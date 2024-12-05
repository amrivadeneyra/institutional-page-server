package dbmodels

import (
	"colegio/server/lib/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name" validate:"required,min=2,max=100"`
	LastName     string             `bson:"last_name" validate:"required,min=2,max=100"`
	Password     string             `bson:"password"  validate:"required,min=6"`
	Email        string             `bson:"email" validate:"required"`
	Active       bool               `bson:"active" validate:"required"`
	Role         string             `bson:"rol" validate:"eq=user|eq=admin|eq=superadmin|eq=sales|eq=service"`
	TokenVersion int64              `bson:"token_version"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	Color        string             `bson:"color"`
	Avatar       string             `bson:"avatar"`
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

type UserByRole struct {
	Role  string `bson:"rol"`
	Users []User `bson:"users"`
}

func (u *UserByRole) ToModel() *models.UserByRole {
	if u == nil {
		return nil
	}

	userByRoleModel := &models.UserByRole{
		Role: u.Role,
	}

	for _, user := range u.Users {
		userByRoleModel.Users = append(userByRoleModel.Users, *user.ToModel())
	}

	return userByRoleModel
}
