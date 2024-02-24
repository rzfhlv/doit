package config

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/utilities/hasher"
	uJwt "github.com/rzfhlv/doit/utilities/jwt"
	logrus "github.com/rzfhlv/doit/utilities/log"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	aJaeger "github.com/rzfhlv/doit/adapter/jaeger"
	aMongo "github.com/rzfhlv/doit/adapter/mongo"
	aPostgres "github.com/rzfhlv/doit/adapter/postgres"
	aRedis "github.com/rzfhlv/doit/adapter/redis"
)

type Config struct {
	Postgres *sqlx.DB
	Mongo    *mongo.Database
	Redis    *redis.Client
	aJaeger.Jaeger
	JWTImpl uJwt.JWTInterface
	Utils   Utils
}

type Utils struct {
	JWTImpl uJwt.JWTInterface
	Hasher  hasher.HashPassword
}

func Init() *Config {
	// load environment
	err := godotenv.Load()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Load Environment, %v", err.Error()))
		os.Exit(1)
	}

	// connect to postgres
	postgres, err := aPostgres.NewPostgres()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Postgres Connection, %v", err.Error()))
		os.Exit(1)
	}

	// connect to mongo
	mongo, err := aMongo.NewMongo()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Mongo Connection, %v", err.Error()))
		os.Exit(1)
	}

	// connect to redis
	redis, err := aRedis.NewRedis()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Redis Connection, %v", err.Error()))
		os.Exit(1)
	}

	// init jaeger
	jaeger, err := aJaeger.NewJaeger()
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Jaeger Connection, %v", err.Error()))
		os.Exit(1)
	}

	jwtImpl := uJwt.JWTImpl{}
	hasher := hasher.HasherPassword{}

	return &Config{
		Postgres: postgres.GetDB(),
		Mongo:    mongo.GetDB(),
		Redis:    redis.GetClient(),
		Jaeger: aJaeger.Jaeger{
			Tracer: jaeger.Tracer,
			Closer: jaeger.Closer,
		},
		JWTImpl: &jwtImpl,
		Utils: Utils{
			JWTImpl: &jwtImpl,
			Hasher:  &hasher,
		},
	}
}
