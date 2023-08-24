package seeders

func (s *Seed) InvestorSeed() {
	for i := 1; i <= 1000; i++ {
		investorQuery := `INSERT INTO investors (name) VALUES ($1);`
		name := s.genrator.GenerateName()
		_, _ = s.cfg.Postgres.Exec(investorQuery, name)
	}
}
