package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateResult struct {
	InsertID   primitive.ObjectID
}
