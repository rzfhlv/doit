package repository

import (
	"context"
	"doit/modules/investor/model"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	SaveMongo(ctx context.Context, investor model.Investor) error
	UpsertMongo(ctx context.Context, investor model.Investor) error
	GetPsql(ctx context.Context) ([]model.Investor, error)
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
		return err
	}
	return nil
}

func (r *Repository) GetPsql(ctx context.Context) ([]model.Investor, error) {
	investors := []model.Investor{}
	err := r.db.Select(&investors, "SELECT * FROM investors")
	if err != nil {
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
		return err
	}
	return nil
}
