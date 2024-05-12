package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	pgdb "gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"github.com/zepif/Test-service/internal/data"
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

func (q *linkQ) Insert(fullURL, shortURL string) error {
    clauses := map[string]interface{}{
        "full_url":  fullURL,
        "short_url": shortURL,
    }

    stmt := sq.Insert(linksTableName).SetMap(clauses)
    err := q.db.Exec(stmt)
    if err != nil {
        return errors.Wrap(err, "failed to insert link into db")
    }

    return nil
}

type link struct {
    FullURL  string `db:"full_url"`
    ShortURL string `db:"short_url"`
}

func (q *linkQ) Get(id int64) (string, string, error) {
    var l link
    err := q.db.Get(&l, q.sql.Select("full_url", "short_url").From(linksTableName).Where(sq.Eq{"id": id}))
    if err == sql.ErrNoRows {
        return "", "", nil
    }
    if err != nil {
        return "", "", errors.Wrap(err, "failed to get link from db")
    }

    return l.FullURL, l.ShortURL, nil
}
