package main

import (
	"doit/config"
	"log"

	"github.com/pressly/goose/v3"
)

func Migrate(cfg *config.Config) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("err goose dialect: %v", err.Error())
	}
	if err := goose.Up(cfg.Postgres.DB, "migrations"); err != nil {
		log.Printf("err migrations: %v", err.Error())
	}
}
