package repository

import (
	"context"

	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/utilities/param"
)

func (r *Repository) GetPsql(ctx context.Context) (investors []model.Investor, err error) {
	err = r.db.Select(&investors, "SELECT * FROM investors")
	return
}

func (r *Repository) GetAll(ctx context.Context, param param.Param) (investors []model.Investor, err error) {
	err = r.db.Select(&investors, `SELECT * FROM investors ORDER BY investors.id DESC LIMIT $1 OFFSET $2;`, param.Limit, param.CalculateOffset())
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (investor model.Investor, err error) {
	err = r.db.Get(&investor, `SELECT * FROM investors WHERE id = $1;`, id)
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, `SELECT count(*) FROM investors;`)
	return
}

func (r *Repository) Generate(ctx context.Context, name string) (err error) {
	_, err = r.db.Exec(`INSERT INTO investors (name) VALUES ($1) RETURNING id;`, name)
	return
}
