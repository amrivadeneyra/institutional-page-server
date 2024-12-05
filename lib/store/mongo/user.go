package mongo

import (
	"colegio/server/lib/mongodb"
	"colegio/server/lib/store/dbmodels"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Store) CreateUser(ctx context.Context, user dbmodels.User) (primitive.ObjectID, error) {
	result, err := mongodb.User().InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if insertedID, ok := result.InsertedID.(primitive.ObjectID); ok {
		return insertedID, nil
	}

	return primitive.NilObjectID, errors.New("can't convert InsertOneResult to primitive.ObjectID")
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*dbmodels.User, error) {
	var user dbmodels.User
	err := mongodb.User().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUser(ctx context.Context, userID primitive.ObjectID) (*dbmodels.User, error) {
	var user dbmodels.User
	err := mongodb.User().FindOne(ctx, bson.M{
		"_id": userID,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
