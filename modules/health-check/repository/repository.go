package repository

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IRepository interface {
	Ping(ctx context.Context) (err error)
	MongoPing(ctx context.Context) (err error)
	RedisPing(ctx context.Context) (err error)
}

type Repository struct {
	db      *sqlx.DB
	dbMongo *mongo.Database
	redis   *redis.Client
}

func NewRepository(db *sqlx.DB, dbMongo *mongo.Database, redis *redis.Client) IRepository {
	return &Repository{
		db:      db,
		dbMongo: dbMongo,
		redis:   redis,
	}
}

func (r *Repository) Ping(ctx context.Context) (err error) {
	err = r.db.Ping()
	if err != nil {
		log.Printf("[ERROR] Health Check Repo Postgres Ping: %v", err.Error())
	}
	return
}

func (r *Repository) MongoPing(ctx context.Context) (err error) {
	err = r.dbMongo.Client().Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("[ERROR] Health Check Repo Mongo Ping: %v", err.Error())
	}
	return
}

func (r *Repository) RedisPing(ctx context.Context) (err error) {
	err = r.redis.Ping(ctx).Err()
	if err != nil {
		log.Printf("[ERROR] Health Check Repo Redis Ping: %v", err.Error())
	}
	return
}
