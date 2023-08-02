package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/database"
	"github.com/rzfhlv/doit/routes"
	"github.com/rzfhlv/doit/service"

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
	e := routes.ListRoutes(svc)

	// start cron job
	// s := gocron.NewScheduler()
	// s.Every(10).Seconds().Do(svc.Investor.MigrateInvestors, context.Background())
	// <-s.Start()

	e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
