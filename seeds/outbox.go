package seeds

import (
	"context"
	"doit/modules/investor/model"
	"time"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Seed) OutboxSeed(ctx context.Context) {
	for i := 1; i <= 10; i++ {
		now := time.Now()
		outbox := model.Outbox{
			Identifier: int64(i),
			Payload:    faker.Email(),
			Event:      "INVESTOR",
			Status:     "PROCESSING",
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		_, _ = s.cfg.Mongo.Collection("outbox").UpdateOne(ctx,
			bson.M{
				"id":    int64(i),
				"event": "INVESTOR",
			},
			bson.M{
				"$set": outbox,
			}, &options.UpdateOptions{
				Upsert: options.Update().SetUpsert(true).Upsert,
			})
	}
}
