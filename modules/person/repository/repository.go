package repository

import (
	"context"
	"doit/modules/person/model"
	"doit/utilities"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	GetAll(ctx context.Context, filter utilities.Param) (persons []model.Person, err error)
	Count(ctx context.Context) (total int64, err error)
}

type Repository struct {
	dbMongo *mongo.Database
}

func NewRepository(dbMongo *mongo.Database) IRepository {
	return &Repository{
		dbMongo: dbMongo,
	}
}

func (r *Repository) GetAll(ctx context.Context, param utilities.Param) (persons []model.Person, err error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})
	findOptions.SetSkip(int64(param.CalculateOffset()))
	findOptions.SetLimit(int64(param.Limit))

	cursor, err := r.dbMongo.Collection("investors").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("[ERROR] Person Repo GetAll: %v", err.Error())
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person model.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	total, err = r.dbMongo.Collection("investors").CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("[ERROR] Person Repo Count: %v", err.Error())
	}
	return
}
