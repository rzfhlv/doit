package database

import (
	"doit/config"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(cfg *config.Config) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("err goose dialect: %v", err.Error())
	}
	if err := goose.Up(cfg.Postgres.DB, "migrations"); err != nil {
		log.Printf("err migrations: %v", err.Error())
	}
}
