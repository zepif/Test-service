package pg

import (
	"github.com/zepif/Test-service/internal/data"
	pgdb "gitlab.com/distributed_lab/kit/pgdb"
)

func NewStorage(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db.Clone(),
	}
}

type masterQ struct {
	db *pgdb.DB
}

func (m *masterQ) New() data.MasterQ {
	return NewStorage(m.db)
}

func (m *masterQ) Link() data.LinkQ {
	return newLinkQ(m.db)
}
