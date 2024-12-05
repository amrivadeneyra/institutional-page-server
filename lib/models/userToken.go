package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserToken struct {
	ID           primitive.ObjectID
	Name         string
	LastName     string
	Email        string
	Rol          string
	Token        string
	RefreshToken string
	Avatar       string
}
