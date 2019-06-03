package service

import (
	"context"

	"github.com/taxio/gitcrow/domain/model"
)

type CloneService interface {
	DelegateToWorker(ctx context.Context, username, saveDir string, repos []*model.GitRepo) error
}

type cloneServiceImpl struct {
	// infra instance
}

func NewCloneService() CloneService {
	return &cloneServiceImpl{}
}

func (s *cloneServiceImpl) DelegateToWorker(ctx context.Context, username, saveDir string, repos []*model.GitRepo) error {
	return nil
}
