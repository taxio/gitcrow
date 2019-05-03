package service

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/taxio/gitcrow/domain/model"
	"golang.org/x/oauth2"
)

type DownloadService interface {
	DelegateToWorker(ctx context.Context, username, saveDir, accessToken string, repos []*model.GitRepo) error
}

type downloadServiceImpl struct {
	// infra instance
}

func NewDownloadService() *downloadServiceImpl {
	return &downloadServiceImpl{}
}

func (s *downloadServiceImpl) DelegateToWorker(ctx context.Context, username, saveDir, accessToken string, repos []*model.GitRepo) error {
	// TODO: すでにリクエストを投げているかどうか確認

	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
	client := github.NewClient(tc)

	// check credential
	_, _, err := client.RateLimits(ctx)
	if err != nil {
		// TODO: handle error (using github.ErrorResponse)
		return err
	}

	// TODO: validate user save directory

	go func(){
		for _, repo := range repos {
			fmt.Println(repo)
			// TODO: Cache存在確認
			// TODO: API経由でダウンロード

			// TODO: DBに記録
			// TODO: Cacheに保存
			// TODO: saveDirに保存
		}
	}()

	return nil
}
