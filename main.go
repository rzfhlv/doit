package main

import (
	"doit/config"
	"doit/service"
	"embed"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
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

	e := echo.New()

	v1 := e.Group("/v1")
	v1.GET("/investor", svc.InvestorHandler.GetAll)
	v1.GET("/investor/:id", svc.InvestorHandler.GetByID)

	// start cron job
	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(svc.Investor.MigrateInvestors, context.Background())
	// <-s.Start()

	e.Start(":8090")
}
