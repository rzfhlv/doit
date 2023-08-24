package seeders

import (
	"context"

	"github.com/rzfhlv/doit/config"
	"github.com/rzfhlv/doit/utilities/faker"
)

type ISeed interface {
	InvestorSeed()
	OutboxSeed(ctx context.Context)
}

type Seed struct {
	cfg      *config.Config
	genrator faker.Generator
}

func NewSeed(cfg *config.Config) ISeed {
	return &Seed{
		cfg:      cfg,
		genrator: &faker.FakerGenerator{},
	}
}
