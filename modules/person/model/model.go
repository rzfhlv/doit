package model

type Person struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}
