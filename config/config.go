package config

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Postgres *sqlx.DB
	Mongo    *mongo.Database
}

func Init() *Config {
	// connect to postgres
	psql := NewPostgres()

	// connect to mongo
	mongo := NewMongo()

	return &Config{
		Postgres: psql,
		Mongo:    mongo,
	}
}
