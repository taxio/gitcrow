package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/app/di"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	ErrAlreadyAcceptedDownloadRequest = errors.New("the user already requested download")
	ErrTagNotFound                    = errors.New("Tag not found")
	ErrGitHubAuth                     = errors.New("github authentication failed")
	ErrPathValidation                 = errors.New("invalid path")
)

type DownloadService interface {
	DelegateToWorker(ctx context.Context, username, projectName, accessToken string, repos []*model.GitRepo) error
}

type downloadServiceImpl struct {
	requestUsers []string // リクエストを投げているuser list

	cacheStore  repository.CacheStore
	recordStore repository.RecordStore
	reportStore repository.ReportStore
	userStore   repository.UserStore

	mu sync.Mutex
}

func NewDownloadService(component di.AppComponent) DownloadService {
	return &downloadServiceImpl{
		requestUsers: []string{},
		cacheStore:   component.CacheStore(),
		recordStore:  component.RecordStore(),
		reportStore:  component.ReportStore(),
		userStore:    component.UserStore(),
	}
}

func (s *downloadServiceImpl) DelegateToWorker(ctx context.Context, username, projectName, accessToken string, repos []*model.GitRepo) error {
	// check whether the user already requested
	hasReq, err := s.alreadyRequested(ctx, username)
	if err != nil {
		return err
	}
	if !hasReq {
		return errors.WithStack(ErrAlreadyAcceptedDownloadRequest)
	}
	err = s.addRequestUser(ctx, username)
	if err != nil {
		return err
	}

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	client := github.NewClient(tc)

	// check credential
	_, _, err = client.RateLimits(ctx)
	if err != nil {
		// TODO: handle error (using github.ErrorResponse)
		_ = s.removeRequestUser(ctx, username)
		grpclog.Errorf("%+v\n", err)
		return errors.WithStack(ErrGitHubAuth)
	}

	// validate user save directory
	err = s.userStore.ValidatePathname(ctx, username, projectName)
	if err != nil {
		grpclog.Errorf("%+v\n", err)
		return errors.WithStack(ErrPathValidation)
	}

	// make user save dir
	err = s.userStore.MakeUserProjectDir(ctx, username, projectName)
	if err != nil {
		return err
	}

	go s.runWorker(client, username, projectName, repos)

	return nil
}

func (s *downloadServiceImpl) runWorker(client *github.Client, username, projectName string, repos []*model.GitRepo) {
	grpclog.Infof("start %s download worker\n", username)
	ctx := context.Background()
	var reports []*model.Report
	for _, repo := range repos {
		filename := fmt.Sprintf("%s-%s-%s.zip", repo.Owner, repo.Repo, repo.Tag)

		// check existence in cache
		var data []byte
		exists, err := s.cacheStore.Exists(ctx, filename)
		if err != nil {
			grpclog.Errorf("%+v\n", err)
		}
		if exists {
			data, err = s.cacheStore.LoadZip(ctx, filename)
			if err != nil {
				grpclog.Errorf("%+v\n", err)
			}
		}

		if !exists || data == nil {
			// download zip data
			data, err = s.downloadRepository(ctx, client, repo)
			if err != nil {
				var msg string
				if errors.Cause(err) == ErrTagNotFound {
					msg = "Tag not found"
				} else {
					msg = "InternalError: Cannot download"
					grpclog.Errorf("%+v\n", err)
				}
				reports = append(reports, &model.Report{
					GitRepo: repo,
					Success: false,
					Message: msg,
				})
				continue
			}

			// record to DB
			exists, err = s.recordStore.Exists(ctx, repo)
			if err != nil {
				grpclog.Errorf("%+v\n", errors.WithStack(err))
			}
			if !exists {
				err = s.recordStore.Insert(ctx, repo)
				if err != nil {
					grpclog.Errorf("db record failed: %+v, %+v\n", repo, err)
				}
			}

			// save to cache dir
			err = s.cacheStore.Save(ctx, filename, data)
			if err != nil {
				grpclog.Errorf("cannot save to cache: %s, %+v\n", filename, err)
			}
		}

		// save to user dir
		err = s.userStore.Save(ctx, username, projectName, filename, data)
		if err != nil {
			grpclog.Errorf("%+v\n", err)
			reports = append(reports, &model.Report{
				GitRepo: repo,
				Success: false,
				Message: "InternalError: Cannot save to user directory",
			})
		} else {
			reports = append(reports, &model.Report{
				GitRepo: repo,
				Success: true,
				Message: "",
			})
		}
	}

	// report to user
	slackId, ok, err := s.recordStore.GetSlackId(ctx, username)
	if err != nil {
		grpclog.Errorf("%+v\n", err)
	}
	if !ok {
		slackId = username
	}
	err = s.reportStore.Notify(ctx, slackId, "finish download worker")
	if err != nil {
		grpclog.Errorf("%+v\n", err)
	}
	err = s.reportStore.ReportToFile(ctx, username, projectName, reports)
	if err != nil {
		grpclog.Errorf("%+v\n", err)
	}

	defer func() {
		err = s.removeRequestUser(ctx, username)
		if err != nil {
			grpclog.Errorf("remove request user failed: #s, %+v\n", username, err)
		}
	}()

	grpclog.Infof("finish %s download worker\n", username)
}

func (s *downloadServiceImpl) alreadyRequested(ctx context.Context, username string) (bool, error) {
	for _, user := range s.requestUsers {
		if user == username {
			return false, nil
		}
	}
	return true, nil
}

func (s *downloadServiceImpl) addRequestUser(ctx context.Context, username string) error {
	s.mu.Lock()
	s.requestUsers = append(s.requestUsers, username)
	s.mu.Unlock()
	return nil
}

func (s *downloadServiceImpl) removeRequestUser(ctx context.Context, username string) error {
	s.mu.Lock()
	for i, user := range s.requestUsers {
		if user == username {
			s.requestUsers = append(s.requestUsers[:i], s.requestUsers[i+1:]...)
		}
	}
	s.mu.Unlock()
	return nil
}

func (s *downloadServiceImpl) downloadRepository(ctx context.Context, client *github.Client, repo *model.GitRepo) ([]byte, error) {
	// wait for API Limit
	// FYI: https://developer.github.com/v3/#rate-limiting
	time.Sleep(1 * time.Second)

	// get tag list
	tags, _, err := client.Repositories.ListTags(ctx, repo.Owner, repo.Repo, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// check tag existence
	var zipUrl string
	for _, tag := range tags {
		if *tag.Name == repo.Tag {
			zipUrl = *tag.ZipballURL
		}
	}
	if len(zipUrl) == 0 {
		return nil, errors.WithStack(ErrTagNotFound)
	}

	// download zip
	resp, err := http.Get(zipUrl)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}
