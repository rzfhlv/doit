package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres() (*Postgres, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
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
