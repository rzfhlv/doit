package main

import (
	"context"
	"doit/config"
	"doit/service"
	"embed"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jasonlvhit/gocron"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Do it start")

	// init config
	cfg := config.Init()

	// migrations
	Migrate(cfg)

	// HandleArgs(db)

	// load service
	svc := service.NewService(cfg)

	// start cron job
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(svc.Investor.MigrateInvestors, context.Background())
	<-s.Start()
}
