package config

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func NewPostgres() *sqlx.DB {
	// connect to postgres
	db, err := sqlx.Open("postgres", "postgres://doit:verysecret@localhost:5434/doit?sslmode=disable")
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("error ping: %v", err.Error())
	}

	return db
}
