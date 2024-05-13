package pg

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/zepif/Test-service/internal/data"
	pgdb "gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const linksTableName = "urlstorage"

type linkQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func newLinkQ(db *pgdb.DB) data.LinkQ {
	return &linkQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

func (q *linkQ) Insert(FullURL, ShortURL string) error {
	clauses := map[string]interface{}{
		"ShortURL": ShortURL,
		"FullURL":  FullURL,
	}

	stmt := sq.Insert(linksTableName).SetMap(clauses)
	err := q.db.Exec(stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert link into db")
	}

	return nil
}

func (q *linkQ) Get(shortURL string) (string, error) {
	var fullURL string
	err := q.db.Get(&fullURL, q.sql.Select("FullURL").From(linksTableName).Where(sq.Eq{"ShortURL": shortURL}))
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", errors.Wrap(err, "failed to get link from db")
	}

	return fullURL, nil
}
