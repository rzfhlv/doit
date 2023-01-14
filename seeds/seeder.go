package seeds

import (
	"github.com/jmoiron/sqlx"
)

type ISeed interface {
	InvestorSeed()
}

type Seed struct {
	db *sqlx.DB
}

func NewSeed(db *sqlx.DB) ISeed {
	return &Seed{
		db: db,
	}
}
