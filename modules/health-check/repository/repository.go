package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
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
	sp, _ := opentracing.StartSpanFromContext(ctx, "Health Check Repository Ping")
	defer sp.Finish()

	err = r.db.Ping()
	return
}

func (r *Repository) MongoPing(ctx context.Context) (err error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "Health Check Repository MongoPing")
	defer sp.Finish()

	err = r.dbMongo.Client().Ping(ctx, readpref.Primary())
	return
}

func (r *Repository) RedisPing(ctx context.Context) (err error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "Health Check Repository RedisPing")
	defer sp.Finish()

	err = r.redis.Ping(ctx).Err()
	return
}
