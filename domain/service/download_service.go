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
)

var (
	ErrAlreadyAcceptedDownloadRequest = errors.New("the user already requested download")
	ErrTagNotFound                    = errors.New("Tag not found")
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
		return errors.WithStack(err)
	}
	if !hasReq {
		return ErrAlreadyAcceptedDownloadRequest
	}
	err = s.addRequestUser(ctx, username)
	if err != nil {
		return errors.WithStack(err)
	}

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	client := github.NewClient(tc)

	// check credential
	_, _, err = client.RateLimits(ctx)
	if err != nil {
		// TODO: handle error (using github.ErrorResponse)
		_ = s.removeRequestUser(ctx, username)
		return errors.WithStack(err)
	}

	// TODO: validate user save directory

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
		exists, err := s.cacheStore.Exists(ctx, filename)
		if err != nil {
			grpclog.Errorln(err)
		}
		if exists {
			// TODO: handle
			grpclog.Infof("%s is already cached.\n", filename)
			continue
		}

		// download zip data
		data, err := s.downloadRepository(ctx, client, repo)
		if err != nil {
			var (
				code model.ReportCode
				msg  string
			)
			if errors.Cause(err) == ErrTagNotFound {
				code = model.ReportInternalErr
				msg = "Tag not found"
			} else {
				code = model.ReportInternalErr
				msg = "Cannot download"
				grpclog.Errorln(err)
			}
			reports = append(reports, &model.Report{
				GitRepo: repo,
				Code:    code,
				Message: msg,
			})
			continue
		}

		// record to DB
		err = s.recordStore.Insert(ctx, repo)
		if err != nil {
			grpclog.Errorf("db record failed: %#v, %#v\n", repo, err)
		}

		// save to cache dir
		err = s.cacheStore.Save(ctx, filename, data)
		if err != nil {
			grpclog.Errorf("cannot save to cache: %s, %#v\n", filename, err)
		}

		// save to user dir
		err = s.userStore.Save(ctx, username, projectName, filename, data)
		if err != nil {
			grpclog.Errorln(err)
			reports = append(reports, &model.Report{
				GitRepo: repo,
				Code:    model.ReportSaveErr,
				Message: "Cannot save",
			})
		} else {
			reports = append(reports, &model.Report{
				GitRepo: repo,
				Code:    model.ReportSuccess,
				Message: "",
			})
		}
	}

	// report to user
	slackId, ok, err := s.recordStore.GetSlackId(ctx, username)
	if err != nil {
		grpclog.Errorln(errors.WithStack(err))
	}
	if !ok {
		slackId = username
	}
	err = s.reportStore.Notify(ctx, slackId, "finish download worker")
	if err != nil {
		grpclog.Error(errors.WithStack(err))
	}
	err = s.reportStore.ReportToFile(ctx, username, projectName, reports)
	if err != nil {
		grpclog.Error(errors.WithStack(err))
	}

	defer func() {
		err = s.removeRequestUser(ctx, username)
		if err != nil {
			grpclog.Errorf("remove request user failed: #s, %#v\n", username, err)
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
