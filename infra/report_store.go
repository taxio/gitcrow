package infra

import (
	"context"
	"github.com/taxio/gitcrow/domain/repository"
)

type reportStoreImpl struct {
	WebHookURL string
	SaveDir    string
}

func NewReportStore(webHookURL, saveDir string) repository.ReportStore {
	return &reportStoreImpl{
		WebHookURL: webHookURL,
		SaveDir:    saveDir,
	}
}

func (s *reportStoreImpl) Notify(ctx context.Context, username, message string) error {
	return nil
}

func (s *reportStoreImpl) Save(ctx context.Context) error {
	return nil
}
