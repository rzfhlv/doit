package main

import (
	"doit/config"
	"doit/database"
	"doit/route"
	"doit/service"
	"flag"

	// "github.com/jasonlvhit/gocron"

	_ "github.com/lib/pq"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// init config
	cfg := config.Init()

	// migrations
	database.Migrate(cfg)

	// seeders
	database.Seed(cfg, args)

	// load service
	svc := service.NewService(cfg)

	// load route
	e := route.ListRoute(svc)

	// start cron job
	// s := gocron.NewScheduler()
	// s.Every(10).Seconds().Do(svc.Investor.MigrateInvestors, context.Background())
	// <-s.Start()

	e.Start(":8090")
}
