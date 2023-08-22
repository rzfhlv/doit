package repository

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/investor/model"
	logrus "github.com/rzfhlv/doit/utilities/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) SaveMongo(ctx context.Context, investor model.Investor) error {
	_, err := r.dbMongo.Collection("investors").InsertOne(ctx, investor)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Investor Repo SaveMongo, %v", err.Error()))
		return err
	}
	return nil
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
