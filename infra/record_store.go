package infra

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/infra/record"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strings"

	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
)

type recordStoreImpl struct {
	db *sql.DB
}

func NewRecordStore(db *sql.DB) repository.RecordStore {
	return &recordStoreImpl{db: db}
}

func (s *recordStoreImpl) Exists(ctx context.Context, repo *model.GitRepo) (bool, error) {
	return false, nil
}

func (s *recordStoreImpl) Insert(ctx context.Context, repo *model.GitRepo) error {
	return nil
}

func (s *recordStoreImpl) Sync(ctx context.Context, repos []*model.GitRepo) error {
	return nil
}

func (s *recordStoreImpl) GetSlackId(ctx context.Context, username string) (string, bool, error) {
	user, err := record.Users(qm.Where("name=?", username)).One(ctx, s.db)
	if err != nil {
		return "", false, errors.Wrap(err, fmt.Sprintf("select %s's slack id", username))
	}
	if user.SlackID.IsZero() {
		return "", false, nil
	}
	slackId := user.SlackID.String
	slackId = strings.TrimSpace(slackId)
	return slackId, true, nil
}