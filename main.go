package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	// start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	if err := cfg.Postgres.Close(); err != nil {
		e.Logger.Fatal(err.Error())
	}

	if err := cfg.Redis.Close(); err != nil {
		e.Logger.Fatal(err.Error())
	}

	if err := cfg.Mongo.Client().Disconnect(context.Background()); err != nil {
		e.Logger.Fatal(err.Error())
	}

	if err := cfg.Jaeger.Closer.Close(); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
