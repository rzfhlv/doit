package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id,omitempty" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	Email     string    `json:"email" db:"email" validate:"required"`
	Username  string    `json:"username" db:"username" validate:"required"`
	Password  string    `json:"password" db:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) HashedPassword() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
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
