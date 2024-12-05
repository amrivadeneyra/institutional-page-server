package jsonmodels

import (
	"colegio/server/lib/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id"`
	Name         string             `json:"name"`
	LastName     string             `json:"last_name"`
	Password     string             `json:"password"`
	Email        string             `json:"email"`
	Active       bool               `json:"active"`
	Role         string             `json:"role"`
	TokenVersion int64              `json:"token_version"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	Avatar       string             `json:"avatar"`
}

func (u *User) FillFromModel(user *models.User) {
	if user == nil {
		return
	}
	u.ID = user.ID
	u.Name = user.Name
	u.LastName = user.LastName
	u.Password = user.Password
	u.Email = user.Email
	u.Active = user.Active
	u.Role = user.Role
	u.TokenVersion = user.TokenVersion
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	u.Avatar = user.Avatar
}

func (u *User) ToModel() *models.User {
	if u == nil {
		return nil
	}
	return &models.User{
		ID:           u.ID,
		Name:         u.Name,
		LastName:     u.LastName,
		Password:     u.Password,
		Email:        u.Email,
		Active:       u.Active,
		Role:         u.Role,
		TokenVersion: u.TokenVersion,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Avatar:       u.Avatar,
	}
}

type VerifyEmail struct {
	Email  string `json:"email"`
	Exists bool   `json:"exists"`
}
