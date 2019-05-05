package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/taxio/gitcrow/domain/model"
	"golang.org/x/oauth2"
	"golang.org/x/xerrors"
)

var ErrAlreadyAcceptedDownloadRequest = xerrors.New("the user already requested download")

type DownloadService interface {
	DelegateToWorker(ctx context.Context, username, saveDir, accessToken string, repos []*model.GitRepo) error
}

type downloadServiceImpl struct {
	requestUsers []string // リクエストを投げているuser list
	// infra instance
}

func NewDownloadService() DownloadService {
	return &downloadServiceImpl{
		requestUsers: []string{},
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

	go func() {
		for _, repo := range repos {
			fmt.Println(repo)
			// TODO: Cache存在確認
			// TODO: API経由でダウンロード

			// TODO: DBに記録
			// TODO: Cacheに保存
			// TODO: saveDirに保存
		}
	}()

	err = s.removeRequestUser(ctx, username)
	if err != nil {
		return err
	}

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
