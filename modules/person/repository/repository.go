package repository

import (
	"context"
	"fmt"

	"github.com/rzfhlv/doit/modules/person/model"
	"github.com/rzfhlv/doit/utilities/param"

	logrus "github.com/rzfhlv/doit/utilities/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	GetAll(ctx context.Context, filter param.Param) (persons []model.Person, err error)
	GetByID(ctx context.Context, id int64) (person model.Person, err error)
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

func (r *Repository) GetAll(ctx context.Context, param param.Param) (persons []model.Person, err error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "id", Value: -1}})
	findOptions.SetSkip(int64(param.CalculateOffset()))
	findOptions.SetLimit(int64(param.Limit))

	cursor, err := r.dbMongo.Collection("investors").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Repo GetAll, %v", err.Error()))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var person model.Person
		err = cursor.Decode(&person)
		if err != nil {
			logrus.Log(nil).Error(fmt.Sprintf("Person Repo Cursor Decode, %v", err.Error()))
			return
		}
		persons = append(persons, person)
	}
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (person model.Person, err error) {
	err = r.dbMongo.Collection("investors").FindOne(ctx, bson.M{"id": id}).Decode(&person)
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Repo GetAll, %v", err.Error()))
	}
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	total, err = r.dbMongo.Collection("investors").CountDocuments(ctx, bson.M{})
	if err != nil {
		logrus.Log(nil).Error(fmt.Sprintf("Person Repo Count, %v", err.Error()))
	}
	return
}
