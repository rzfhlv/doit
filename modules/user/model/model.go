package model

import (
	"time"
)

type User struct {
	ID        int64     `json:"id,omitempty" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required"`
	Username  string    `json:"username" db:"username" validate:"required"`
	Password  string    `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Login struct {
	Username string `json:"username" db:"username" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}

type JWT struct {
	Token   string `json:"token"`
	Expired string `json:"expired"`
}

type Validate struct {
	Token string `json:"token" validate:"required"`
}
