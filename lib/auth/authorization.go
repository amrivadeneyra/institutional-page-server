package auth

import (
	"colegio/server/common/utils"
	"context"

	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DecorateContext(ctx context.Context, t jwt.Token) (context.Context, error) {
	claims, err := t.AsMap(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userRole, ok := claims["role"].(string)

	userIDString, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid user in claims")
	}

	userID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, errors.New("invalid user id in claims")
	}

	ctx = context.WithValue(ctx, utils.UserRoleKey, userRole)
	ctx = context.WithValue(ctx, utils.UserIDKey, userID)
	ctx = context.WithValue(ctx, utils.TokenKey, t)
	return ctx, nil
}

func GetUserRoleFromContext(ctx context.Context) (string, error) {
	value := ctx.Value(utils.UserRoleKey)
	if value == nil {
		return "", nil
	}

	userRole, ok := value.(string)
	if !ok {
		return "", errors.Errorf("cannot convert %v to string", value)
	}
	if userRole == "" {
		return "", errors.New("user role in context is empty")
	}
	return userRole, nil
}

func GetUserIDFromContext(ctx context.Context) (userID primitive.ObjectID, err error) {
	value := ctx.Value(utils.UserIDKey)
	if value == nil {
		return primitive.NilObjectID, nil
	}

	userID, ok := value.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.Errorf("cannot convert %v to primitive.ObjectId", value)
	}
	if userID == primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("user id in context is nil")
	}
	return userID, nil
}

func GetTokenFromContext(ctx context.Context) (token jwt.Token, found bool, err error) {
	value := ctx.Value(utils.TokenKey)
	if value == nil {
		return jwt.New(), false, nil
	}
	token, ok := value.(jwt.Token)
	if !ok {
		return jwt.New(), false, errors.Errorf("cannot convert %v to jwt.Token", value)
	}
	if token == nil {
		return nil, false, errors.New("token in context is nil")
	}
	return token, true, nil
}

func GetTokenVersion(ctx context.Context, t jwt.Token) (int64, error) {

	claims, err := t.AsMap(ctx)
	if err != nil {
		return -1, errors.WithStack(err)
	}

	tokenVersion, ok := claims["token_version"].(float64)
	if !ok {
		return -1, errors.New("invalid token version in claims")
	}
	return int64(tokenVersion), nil
}
