package model

import "time"

type Investor struct {
	ID   int64  `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Outbox struct {
	Identifier int64     `json:"identifier" bson:"id"`
	Payload    string    `json:"payload" bson:"payload"`
	Event      string    `json:"event" bson:"event"`
	Status     string    `json:"status" bson:"status"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}
