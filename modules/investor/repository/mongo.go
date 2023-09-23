package repository

import (
	"context"

	"github.com/rzfhlv/doit/modules/investor/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	INVESTORS = "investors"
	OUTBOX    = "outbox"
)

func (r *Repository) SaveMongo(ctx context.Context, investor model.Investor) (err error) {
	_, err = r.dbMongo.Collection(INVESTORS).InsertOne(ctx, investor)
	return
}

func (r *Repository) UpsertMongo(ctx context.Context, investor model.Investor) (err error) {
	_, err = r.dbMongo.Collection(INVESTORS).
		UpdateOne(ctx,
			bson.M{
				"id": investor.ID,
			},
			bson.M{
				"$set": investor,
			}, &options.UpdateOptions{
				Upsert: options.Update().SetUpsert(true).Upsert,
			})
	return
}

func (r *Repository) UpsertOutbox(ctx context.Context, outbox model.Outbox) (err error) {
	_, err = r.dbMongo.Collection(OUTBOX).
		UpdateOne(ctx,
			bson.M{
				"identifier": outbox.Identifier,
			},
			bson.M{
				"$set": outbox,
			}, &options.UpdateOptions{
				Upsert: options.Update().SetUpsert(true).Upsert,
			})
	return
}

func (r *Repository) DeleteOutbox(ctx context.Context, identifier int64) (err error) {
	_, err = r.dbMongo.Collection(OUTBOX).DeleteOne(ctx, bson.M{"identifier": identifier})
	return
}
