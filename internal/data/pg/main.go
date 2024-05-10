package pg

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/zepif/Test-service/internal/data"
)

func NewStorage(db *pgdb) data.URLStorage {
    return &URLStorage {
        db: db.Clone(),
    }
}

type URLStorage struct {
    db *pgdb.DB
}

func (m *URLStorage) New() data.URLStorage {
    return NewStorage(m.db)
}

func (m *URLStorage) Link() data.LinkQ {
	return newLinkQ(m.db)
}


