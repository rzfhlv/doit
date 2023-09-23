package usecase

import (
	"context"

	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/modules/investor/repository"
	"github.com/rzfhlv/doit/utilities/param"

	"github.com/bxcodec/faker/v3"
)

type IUsecase interface {
	MigrateInvestors(ctx context.Context) (err error)
	ConventionalMigrate(ctx context.Context) (err error)
	GetAll(ctx context.Context, param *param.Param) (investors []model.Investor, err error)
	GetByID(ctx context.Context, id int64) (investor model.Investor, err error)
	Generate(ctx context.Context) (err error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetAll(ctx context.Context, param *param.Param) (investors []model.Investor, err error) {
	investors, err = u.repo.GetAll(ctx, *param)
	if err != nil {
		return
	}
	if len(investors) < 1 {
		investors = []model.Investor{}
	}
	total, err := u.repo.Count(ctx)
	param.Total = total
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (investor model.Investor, err error) {
	investor, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Generate(ctx context.Context) (err error) {
	for i := 1; i <= 1000; i++ {
		name := faker.Name()
		err = u.repo.Generate(ctx, name)
		if err != nil {
			break
		}
	}
	return
}
