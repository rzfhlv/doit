package mongo

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
	mongoError  error
)

type Mongo struct {
	client *mongo.Client
}

func NewMongo() (*Mongo, error) {
	mongoOnce.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))
		clientOptions := options.Client()
		clientOptions.ApplyURI(uri)

		var err error
		mongoClient, err = mongo.NewClient(clientOptions)
		if err != nil {
			mongoError = err
		}

		err = mongoClient.Connect(context.Background())
		if err != nil {
			mongoError = err
		}
	})

	if mongoError != nil {
		return nil, mongoError
	}

	return &Mongo{
		client: mongoClient,
	}, nil
}

func (m *Mongo) GetDB() *mongo.Database {
	return m.client.Database(os.Getenv("MONGO_DB"))
}

func (m *Mongo) GetClient() *mongo.Client {
	return m.client
}

func (m *Mongo) Close() {
	if m.client != nil {
		m.client.Disconnect(context.Background())
	}
}
