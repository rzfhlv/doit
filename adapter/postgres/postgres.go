package postgres

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	psqlDB    *sqlx.DB
	psqlOnce  sync.Once
	psqlError error
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres() (*Postgres, error) {
	psqlOnce.Do(func() {
		var err error
		psqlDB, err = sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
		if err != nil {
			psqlError = err
		}

		err = psqlDB.Ping()
		if err != nil {
			psqlError = err
		}
	})

	if psqlError != nil {
		return nil, psqlError
	}

	return &Postgres{
		db: psqlDB,
	}, nil
}

func (p *Postgres) GetDB() *sqlx.DB {
	return p.db
}

func (p *Postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
