package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rzfhlv/doit/modules/person/model"
	"github.com/rzfhlv/doit/modules/person/repository"
	"github.com/rzfhlv/doit/utilities/param"
)

type IUsecase interface {
	GetAll(ctx context.Context, param *param.Param) (persons []model.Person, err error)
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

func (u *Usecase) GetAll(ctx context.Context, param *param.Param) (persons []model.Person, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Person Usecase GetAll")
	defer sp.Finish()

	persons, err = u.repo.GetAll(ctx, *param)
	if err != nil {
		return
	}
	if len(persons) < 1 {
		persons = []model.Person{}
	}

	total, err := u.repo.Count(ctx)
	param.Total = total
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (person model.Person, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Person Usecase GetByID")
	defer sp.Finish()

	person, err = u.repo.GetByID(ctx, id)
	return
}
