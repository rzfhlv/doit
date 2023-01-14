package main

import (
	"doit/seeds"
	"flag"
	"fmt"
	"log"
	"os"

	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	fmt.Println("Hello doit")

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
