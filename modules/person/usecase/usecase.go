package usecase

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/person/model"
	"github.com/rzfhlv/doit/modules/person/repository"
	"github.com/rzfhlv/doit/utilities"
	logrus "github.com/rzfhlv/doit/utilities/log"
)

type IUsecase interface {
	GetAll(ctx context.Context, param *utilities.Param) (persons []model.Person, err error)
	GetByID(ctx context.Context, id int64) (person model.Person, err error)
}

type Usecase struct {
	repo repository.IRepository
}

func NewUsecase(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetAll(ctx context.Context, param *utilities.Param) (persons []model.Person, err error) {
	persons, err = u.repo.GetAll(ctx, *param)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Usecase GetAll, %v", err.Error()))
		return
	}
	if len(persons) < 1 {
		persons = []model.Person{}
	}

	total, err := u.repo.Count(ctx)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Usecase Count, %v", err.Error()))
	}
	param.Total = total
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (person model.Person, err error) {
	person, err = u.repo.GetByID(ctx, id)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Usecase GetByID, %v", err.Error()))
	}
	return
}
