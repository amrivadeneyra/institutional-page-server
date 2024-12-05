package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID
	Name         string
	LastName     string
	Password     string
	Email        string
	Active       bool
	Role         string
	TokenVersion int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Avatar       string
}

type UserByRole struct {
	Role  string
	Users []User
}
