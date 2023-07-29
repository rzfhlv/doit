package seeders

import (
	"context"

	"github.com/rzfhlv/doit/config"
)

type ISeed interface {
	InvestorSeed()
	OutboxSeed(ctx context.Context)
}

type Seed struct {
	cfg *config.Config
}

func NewSeed(cfg *config.Config) ISeed {
	return &Seed{
		cfg: cfg,
	}
}
