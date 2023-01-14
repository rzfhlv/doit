package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() (*mongo.Database, error) {
	// connect to mongo
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://citizix:S3cret@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("doit"), nil
}
