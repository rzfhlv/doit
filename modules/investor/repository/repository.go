package repository

import (
	"context"

	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/utilities/param"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository interface {
	SaveMongo(ctx context.Context, investor model.Investor) error
	UpsertMongo(ctx context.Context, investor model.Investor) error
	UpsertOutbox(ctx context.Context, outbox model.Outbox) error
	DeleteOutbox(ctx context.Context, identifier int64) error
	GetPsql(ctx context.Context) ([]model.Investor, error)
	GetAll(ctx context.Context, param param.Param) (investors []model.Investor, err error)
	GetByID(ctx context.Context, id int64) (investor model.Investor, err error)
	Count(ctx context.Context) (total int64, err error)
	Generate(ctx context.Context, name string) (err error)
}

type Repository struct {
	db      *sqlx.DB
	dbMongo *mongo.Database
}

func NewRepository(db *sqlx.DB, dbMongo *mongo.Database) IRepository {
	return &Repository{
		db:      db,
		dbMongo: dbMongo,
	}
}
