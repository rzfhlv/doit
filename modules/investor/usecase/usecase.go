package usecase

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/modules/investor/repository"
	"github.com/rzfhlv/doit/utilities/param"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/bxcodec/faker/v3"
)

type IUsecase interface {
	MigrateInvestors(ctx context.Context) error
	ConventionalMigrate(ctx context.Context) error
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
		logrus.Log(nil).Error(fmt.Sprintf("Investor Usecase GetAll, %v", err.Error()))
		return
	}
	if len(investors) < 1 {
		investors = []model.Investor{}
	}
	total, err := u.repo.Count(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Usecase GetAll Count, %v", err.Error()))
	}
	param.Total = total
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (investor model.Investor, err error) {
	investor, err = u.repo.GetByID(ctx, id)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Usecase GetByID, %v", err.Error()))
	}
	return
}

func (u *Usecase) Generate(ctx context.Context) (err error) {
	for i := 1; i <= 1000; i++ {
		name := faker.Name()
		err = u.repo.Generate(ctx, name)
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Investor Usecase Generate, %v", err.Error()))
			break
		}
	}
	return
}
