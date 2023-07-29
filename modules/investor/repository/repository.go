package repository

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/utilities"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	SaveMongo(ctx context.Context, investor model.Investor) error
	UpsertMongo(ctx context.Context, investor model.Investor) error
	UpsertOutbox(ctx context.Context, outbox model.Outbox) error
	DeleteOutbox(ctx context.Context, identifier int64) error
	GetPsql(ctx context.Context) ([]model.Investor, error)
	GetAll(ctx context.Context, param utilities.Param) (investors []model.Investor, err error)
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

func (r *Repository) SaveMongo(ctx context.Context, investor model.Investor) error {
	_, err := r.dbMongo.Collection("investors").InsertOne(ctx, investor)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo SaveMongo, %v", err.Error()))
		return err
	}
	return nil
}

func (r *Repository) GetPsql(ctx context.Context) ([]model.Investor, error) {
	investors := []model.Investor{}
	err := r.db.Select(&investors, "SELECT * FROM investors")
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo GetPsql, %v", err.Error()))
		return nil, err
	}
	return investors, nil
}

func (r *Repository) UpsertMongo(ctx context.Context, investor model.Investor) error {
	_, err := r.dbMongo.Collection("investors").
		UpdateOne(ctx,
			bson.M{
				"id": investor.ID,
			},
			bson.M{
				"$set": investor,
			}, &options.UpdateOptions{
				Upsert: options.Update().SetUpsert(true).Upsert,
			})
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo UpsertMongo, %v", err.Error()))
		return err
	}
	return nil
}

func (r *Repository) UpsertOutbox(ctx context.Context, outbox model.Outbox) error {
	_, err := r.dbMongo.Collection("outbox").
		UpdateOne(ctx,
			bson.M{
				"identifier": outbox.Identifier,
			},
			bson.M{
				"$set": outbox,
			}, &options.UpdateOptions{
				Upsert: options.Update().SetUpsert(true).Upsert,
			})
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo UpsertOutbox, %v", err.Error()))
		return err
	}
	return nil
}

func (r *Repository) DeleteOutbox(ctx context.Context, identifier int64) error {
	_, err := r.dbMongo.Collection("outbox").DeleteOne(ctx, bson.M{"identifier": identifier})
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo DeleteOutbox, %v", err.Error()))
		return err
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context, param utilities.Param) (investors []model.Investor, err error) {
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
