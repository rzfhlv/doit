package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() (*mongo.Database, error) {
	// connect to mongo
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
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
