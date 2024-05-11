package pg

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/zepif/Test-service/internal/data"
)

const linksTableName = "URLStorage"

type LinkQ interface {
	Insert(fullURL, shortURL string) error
	Get(shortURL string) (string, error)
}

type linkQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func newLinkQ(db *pgdb.DB) LinkQ {
	return &linkQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

func (q *linkQ) Insert(fullURL, shortURL string) error {
	clauses := map[string]interface{}{
		"full_url":  fullURL,
		"short_url": shortURL,
	}
    stmt := sq.Insert(linksTableName).SetMap(clauses)

	_, err := q.db.Exec(stmt)
	if err != nil {
		return errors.Wrap(err, "failed to insert link into db")
	}

	return nil
}

func (q *linkQ) Get(id int64) (string, string, error) {
	var fullURL, shortURL string
	err := q.db.Get(&fullURL, &shortURL, q.sql.Select("full_url", "short_url").From(linksTableName).Where(sq.Eq{"id": id}))
	if err == sql.ErrNoRows {
		return "", "", nil
	}
	if err != nil {
		return "", "", errors.Wrap(err, "failed to get link from db")
	}

	return fullURL, shortURL, nil
}
