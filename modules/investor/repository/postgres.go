package repository

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/investor/model"
	logrus "github.com/rzfhlv/doit/utilities/log"
	"github.com/rzfhlv/doit/utilities/param"
)

func (r *Repository) GetPsql(ctx context.Context) ([]model.Investor, error) {
	investors := []model.Investor{}
	err := r.db.Select(&investors, "SELECT * FROM investors")
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo GetPsql, %v", err.Error()))
		return nil, err
	}
	return investors, nil
}

func (r *Repository) GetAll(ctx context.Context, param param.Param) (investors []model.Investor, err error) {
	err = r.db.Select(&investors, `SELECT * FROM investors ORDER BY investors.id DESC LIMIT $1 OFFSET $2;`, param.Limit, param.CalculateOffset())
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo GetAll, %v", err.Error()))
	}
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (investor model.Investor, err error) {
	err = r.db.Get(&investor, `SELECT * FROM investors WHERE id = $1;`, id)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo GetByID, %v", err.Error()))
	}
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, `SELECT count(*) FROM investors;`)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo Count, %v", err.Error()))
	}
	return
}

func (r *Repository) Generate(ctx context.Context, name string) (err error) {
	_, err = r.db.Exec(`INSERT INTO investors (name) VALUES ($1) RETURNING id;`, name)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo Generate, %v", err.Error()))
	}
	return
}