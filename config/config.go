package config

import (
	"fmt"
	"os"

	logrus "doit/utilities/log"

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
		logrus.Log(nil).Error(fmt.Sprintf("Load Environment, %v", err.Error()))
		os.Exit(1)
	}

	// connect to postgres
	psql, err := NewPostgres()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Psql Connection, %v", err.Error()))
		os.Exit(1)
	}

	// connect to mongo
	mongo, err := NewMongo()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Mongo Connection, %v", err.Error()))
		os.Exit(1)
	}

	// connect to redis
	redis, err := NewRedis()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Redis Connection, %v", err.Error()))
		os.Exit(1)
	}

	return &Config{
		Postgres: psql,
		Mongo:    mongo,
		Redis:    redis,
	}
}
