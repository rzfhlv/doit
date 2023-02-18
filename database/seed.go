package database

import (
	"doit/config"
	"doit/database/seeders"
)

func Seed(cfg *config.Config, args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "seed":

			seed := seeders.NewSeed(cfg)
			seed.InvestorSeed()
			// seed.OutboxSeed(context.Background())
		}
	}
}
