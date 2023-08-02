package mongo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
}

func NewMongo() (*Mongo, error) {
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

	return &Mongo{
		client: client,
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
