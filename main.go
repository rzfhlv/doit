package main

import (
	"context"
	"doit/modules/investor/repository"
	"doit/modules/investor/usecase"
	"doit/seeds"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Do it start")

	// connect to postgres
	db, err := sqlx.Open("postgres", "postgres://doit:verysecret@localhost:5434/doit?sslmode=disable")
	if err != nil {
		log.Printf("error: %v", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("error ping: %v", err.Error())
	}

	// migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Printf("err goose dialect: %v", err.Error())
	}
	if err := goose.Up(db.DB, "migrations"); err != nil {
		log.Printf("err migrations: %v", err.Error())
	}

	// handleArgs(db)

	// connect to mongo
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://citizix:S3cret@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Printf("err mongo: %v", err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("err mongo connect: %v", err.Error())
	}

	investorRepo := repository.NewRepository(db, client)
	investorUsecase := usecase.NewUsecase(investorRepo)
	start := time.Now()
	err = investorUsecase.MigrateInvestors(context.Background())
	if err != nil {
		log.Printf("error migration: %v\n", err.Error())
	}
	duration := time.Since(start)

	fmt.Println("Waaaaack!")
	fmt.Printf("Done in %v seconds\n", duration.Seconds())
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
