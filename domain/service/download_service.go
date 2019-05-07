package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/taxio/gitcrow/app/di"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
)

var ErrAlreadyAcceptedDownloadRequest = xerrors.New("the user already requested download")

type DownloadService interface {
	DelegateToWorker(ctx context.Context, username, saveDir, accessToken string, repos []*model.GitRepo) error
}

type downloadServiceImpl struct {
	requestUsers []string // リクエストを投げているuser list

	cacheStore  repository.CacheStore
	recordStore repository.RecordStore
	reportStore repository.ReportStore
	userStore   repository.UserStore
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

func (s *downloadServiceImpl) DelegateToWorker(ctx context.Context, username, saveDir, accessToken string, repos []*model.GitRepo) error {
	// check whether the user already requested
	hasReq, err := s.alreadyRequested(ctx, username)
	if err != nil {
		return err
	}
	if !hasReq {
		return ErrAlreadyAcceptedDownloadRequest
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
		return err
	}

	// TODO: validate user save directory

	go func(client *github.Client, username, saveDir string, repos []*model.GitRepo) {
		grpclog.Infoln("start %s download worker\n", username)
		ctx := context.Background()
		for _, repo := range repos {
			// check existence in cache
			exists, err := s.cacheStore.Exists(ctx, repo)
			if err != nil {
				// TODO: report
				fmt.Println(err)
				continue
			}
			if exists {
				continue
			}

			// download zip data
			data, err := s.downloadRepository(ctx, client, repo)
			if err != nil {
				// TODO: report
				fmt.Println(err)
				continue
			}
			filename := fmt.Sprintf("%s-%s-%s.zip", repo.Owner, repo.Repo, repo.Tag)

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

			// TODO: saveDirに保存
			err = s.userStore.Save(ctx, filename, data)
			if err != nil {
				// TODO: report
				fmt.Println(err)
			}

			grpclog.Infof("finish %s download worker\n", username)
		}

		// TODO: ユーザーに通知 & report作成

		err = s.removeRequestUser(ctx, username)
		if err != nil {
			grpclog.Errorf("remove request user failed: #s, %#v\n", username, err)
		}
	}(client, username, saveDir, repos)

	return nil
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
	s.requestUsers = append(s.requestUsers, username)
	return nil
}

func (s *downloadServiceImpl) removeRequestUser(ctx context.Context, username string) error {
	for i, user := range s.requestUsers {
		if user == username {
			s.requestUsers = append(s.requestUsers[:i], s.requestUsers[i+1:]...)
		}
	}
	return nil
}

func (s *downloadServiceImpl) downloadRepository(ctx context.Context, client *github.Client, repo *model.GitRepo) ([]byte, error) {
	// get tag list
	tags, _, err := client.Repositories.ListTags(ctx, repo.Owner, repo.Repo, nil)
	if err != nil {
		return nil, err
	}

	// check tag existence
	var zipUrl string
	for _, tag := range tags {
		if *tag.Name == repo.Tag {
			zipUrl = *tag.ZipballURL
		}
	}
	if len(zipUrl) == 0 {
		// TODO: create error for handling
		return nil, fmt.Errorf("tag not found")
	}

	// download zip
	resp, err := http.Get(zipUrl)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
