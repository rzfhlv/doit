package database

import (
	"doit/config"
	"embed"
	"fmt"

	logrus "doit/utilities/log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(cfg *config.Config) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Goose Dialect, %v", err.Error()))
	}
	if err := goose.Up(cfg.Postgres.DB, "migrations"); err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Migrations, %v", err.Error()))
	}
}
