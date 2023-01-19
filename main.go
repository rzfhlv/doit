package main

import (
	"doit/config"
	"doit/route"
	"doit/service"
	"embed"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// init config
	cfg := config.Init()

	// migrations
	Migrate(cfg)

	// HandleArgs(db)

	// load service
	svc := service.NewService(cfg)

	// load route
	e := route.ListRoute(svc)

	// start cron job
	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(svc.Investor.MigrateInvestors, context.Background())
	// <-s.Start()

	e.Start(":8090")
}
