package infra

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/taxio/gitcrow/domain/model"
	"github.com/taxio/gitcrow/domain/repository"
	"github.com/taxio/gitcrow/infra/record"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strings"
)

type recordStoreImpl struct {
	db *sql.DB
}

func NewRecordStore(db *sql.DB) repository.RecordStore {
	return &recordStoreImpl{db: db}
}

func (s *recordStoreImpl) Exists(ctx context.Context, repo *model.GitRepo) (bool, error) {
	cnt, err := record.Cacheds(
		qm.Where("owner=?", repo.Owner),
		qm.Where("repo=?", repo.Repo),
		qm.Where("tag=?", repo.Tag),
	).Count(ctx, s.db)
	if err != nil {
		return false, errors.WithStack(err)
	}

	if cnt == 0 {
		return false, nil
	}
	return true, nil
}

func (s *recordStoreImpl) Insert(ctx context.Context, repo *model.GitRepo) error {
	r := record.Cached{
		Owner: repo.Owner,
		Repo:  repo.Repo,
		Tag:   null.StringFrom(repo.Tag),
	}
	err := r.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		return errors.WithStack(err)
	}

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
