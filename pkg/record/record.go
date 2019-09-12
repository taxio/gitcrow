package record

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/afero"
)

var schema = `
create table repo
(
    id integer not null
        constraint repo_pk
            primary key autoincrement,
    host text not null,
    owner text not null,
    name text not null
);

create unique index repo_host_owner_name_uindex
    on repo (host, owner, name);
`

func CreateDatabase(dbPath string) (*RecordStore, error) {
	r := NewRecordStore(afero.NewOsFs(), dbPath)
	_, err := r.db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func NewRecordStore(fs afero.Fs, dbPath string) *RecordStore {
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil
	}
	return &RecordStore{
		db: db,
		fs: fs,
	}
}

type RecordStore struct {
	db *sqlx.DB
	fs afero.Fs
}

func (r *RecordStore) CreateDatabase(dbPath string) error {
	_, err := r.db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func (r *RecordStore) Add(host, owner, name string) error {
	_, err := r.db.Exec("insert into repo (host, owner, name) VALUES (?, ?, ?)", host, owner, name)
	if err != nil {
		return err
	}
	return nil
}
