package worker

import (
	"context"
	"github.com/k0kubun/pp"
)

type RepositoryInfo struct {
	Owner   string
	Repo    string
	Tag     string
	IsClone bool
}

type RepositorySaveWorker struct {
	queue chan RepositoryInfo
}

func NewRepositorySaveWorker(queue chan RepositoryInfo) (*RepositorySaveWorker, error) {
	return &RepositorySaveWorker{queue: queue}, nil
}

func (w *RepositorySaveWorker) Run(ctx context.Context) error {
	for r := range w.queue {
		pp.Println(r)
	}
	return nil
}
