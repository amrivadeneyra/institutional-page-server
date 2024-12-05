package jsonmodels

import (
	"colegio/server/lib/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserToken struct {
	ID           primitive.ObjectID `json:"_id"`
	Name         string             `json:"name"`
	LastName     string             `json:"last_name"`
	Email        string             `json:"email"`
	Rol          string             `json:"rol"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	Avatar       string             `json:"avatar"`
}

func (u *UserToken) FillFromModel(token *models.UserToken) {
	if token == nil {
		return
	}

	u.ID = token.ID
	u.Name = token.Name
	u.LastName = token.LastName
	u.Email = token.Email
	u.Rol = token.Rol
	u.Token = token.Token
	u.RefreshToken = token.RefreshToken
	u.Avatar = token.Avatar
}
