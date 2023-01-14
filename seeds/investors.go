package seeds

import (
	"github.com/bxcodec/faker/v3"
)

func (s *Seed) InvestorSeed() {
	for i := 1; i <= 1000; i++ {
		investorQuery := `INSERT INTO investors (id, name) VALUES ($1, $2);`
		name := faker.Name()
		_ = s.db.MustExec(investorQuery, i, name)
	}
}
