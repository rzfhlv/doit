package config

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgres() (*sqlx.DB, error) {
	// connect to postgres
	db, err := sqlx.Open("postgres", "postgres://doit:verysecret@localhost:5434/doit?sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
