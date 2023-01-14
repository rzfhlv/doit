package main

import (
	"doit/seeds"
	"flag"
	"os"

	"github.com/jmoiron/sqlx"
)

func HandleArgs(db *sqlx.DB) {
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
