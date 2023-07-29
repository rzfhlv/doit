package database

import (
	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/database/seeders"
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
