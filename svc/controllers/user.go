package controllers

import (
	"colegio/server/common/httpresponses"
	"colegio/server/lib/models"
	"colegio/server/lib/store"
	"colegio/server/lib/store/dbmodels"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, user models.User) (*models.CreateResult, []*httpresponses.ValidationError, error) {
	dbStore := store.NewStoreDefault()
	return createUser(ctx, dbStore, user)
}

func createUser(ctx context.Context, dbStore store.Store, user models.User) (*models.CreateResult, []*httpresponses.ValidationError, error) {
	var err error
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.TokenVersion = 0

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}
	password := string(hash)
	user.Password = password

	dbUser := dbmodels.User{}
	dbUser.FillFromModel(&user)

	insertID, err := dbStore.CreateUser(ctx, dbUser)
	if err != nil {
		return nil, nil, err
	}

	return &models.CreateResult{
		InsertID: insertID,
	}, nil, nil
}

func VerifyEmail(ctx context.Context, email string) (bool, error) {
	return verifyEmail(ctx, store.NewStoreDefault(), email)
}

func verifyEmail(ctx context.Context, dbStore store.Store, email string) (bool, error) {
	user, err := dbStore.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil
}
