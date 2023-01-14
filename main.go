package main

import (
	"context"
	"doit/config"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"
	"doit/seeds"
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/jasonlvhit/gocron"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Do it start")

	cfg := config.Init()

	// migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("err goose dialect: %v", err.Error())
	}
	if err := goose.Up(cfg.Postgres.DB, "migrations"); err != nil {
		log.Printf("err migrations: %v", err.Error())
	}

	// handleArgs(db)

	investorRepo := repository.NewRepository(cfg.Postgres, cfg.Mongo)
	investorUsecase := usecase.NewUsecase(investorRepo)
	start := time.Now()
	duration := time.Since(start)
	fmt.Printf("Done in %v seconds\n", duration.Seconds())

	// start cron job
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(investorUsecase.MigrateInvestors, context.Background())
	<-s.Start()
}

func handleArgs(db *sqlx.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":

			seed := seeds.NewSeed(db)
			seed.InvestorSeed()
			os.Exit(0)
		}
	}
}
