package store

import (
	"colegio/server/lib/store/dbmodels"
	"colegio/server/lib/store/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Store interface {
	CreateUser(ctx context.Context, user dbmodels.User) (primitive.ObjectID, error)
	GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error)
	GetUser(ctx context.Context, userID primitive.ObjectID) (*dbmodels.User, error)

}

func NewStoreDefault() Store {
	return mongo.NewStore()
}
