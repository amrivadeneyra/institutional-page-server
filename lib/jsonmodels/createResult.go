package jsonmodels

import (
	"colegio/server/lib/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateResult struct {
	InsertID primitive.ObjectID `json:"InsertedID"`
}

func (c *CreateResult) FillFromModel(createResult *models.CreateResult) {
	if createResult == nil {
		return
	}

	c.InsertID = createResult.InsertID
}
