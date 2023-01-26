package main

import (
	"context"
	"doit/config"
	"doit/seeds"
	"flag"
	"os"
)

func HandleArgs(cfg *config.Config) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":

			seed := seeds.NewSeed(cfg)
			seed.InvestorSeed()
			seed.OutboxSeed(context.Background())
			os.Exit(0)
		}
	}
}
