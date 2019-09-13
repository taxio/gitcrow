package cloneservice

import "github.com/taxio/gitcrow/pkg/record"

type CloneService struct {
	store *record.RecordStore
}

func NewCloneService(r *record.RecordStore) *CloneService {
	return &CloneService{
		store: r,
	}
}
