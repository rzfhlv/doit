package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() *mongo.Database {
	// connect to mongo
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://citizix:S3cret@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Printf("err mongo: %v", err.Error())
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("err mongo connect: %v", err.Error())
	}

	return client.Database("doit")
}
