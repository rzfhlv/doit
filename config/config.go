package config

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Postgres *sqlx.DB
	Mongo    *mongo.Database
	Redis    *redis.Client
}

func Init() *Config {
	// load environment
	err := godotenv.Load()
	if err != nil {
		log.Printf("error load environment: %v", err.Error())
		os.Exit(1)
	}

	// connect to postgres
	psql, err := NewPostgres()
	if err != nil {
		log.Printf("error psql connection: %v", err.Error())
		os.Exit(1)
	}

	// connect to mongo
	mongo, err := NewMongo()
	if err != nil {
		log.Printf("error mongo connection: %v", err.Error())
		os.Exit(1)
	}

	// connect to redis
	redis, err := NewRedis()
	if err != nil {
		log.Printf("error redis connection: %v", err.Error())
		os.Exit(1)
	}

	return &Config{
		Postgres: psql,
		Mongo:    mongo,
		Redis:    redis,
	}
}
