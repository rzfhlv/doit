package seeders

import (
	"github.com/bxcodec/faker/v3"
)

func (s *Seed) InvestorSeed() {
	for i := 1; i <= 1000; i++ {
		investorQuery := `INSERT INTO investors (name) VALUES ($1);`
		name := faker.Name()
		_ = s.cfg.Postgres.MustExec(investorQuery, name)
	}
}
