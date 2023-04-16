package usecase

import (
	"context"
	"doit/modules/person/model"
	"doit/modules/person/repository"
	"doit/utilities"
	"log"
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
		log.Printf("[ERROR] Person Usecase GetAll: %v", err.Error())
		return
	}
	if len(persons) < 1 {
		persons = []model.Person{}
	}

	total, err := u.repo.Count(ctx)
	if err != nil {
		log.Printf("[ERROR] Person Usecase Count: %v", err.Error())
	}
	param.Total = total
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (person model.Person, err error) {
	person, err = u.repo.GetByID(ctx, id)
	if err != nil {
		log.Printf("[ERROR] Person Usecase GetByID: %v", err.Error())
	}
	return
}
