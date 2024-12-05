package controllers

import (
	"context"
	"colegio/server/common/httpresponses"
	"colegio/server/lib/auth"
	"colegio/server/lib/auth/selfauthorizer"
	"colegio/server/lib/models"
	"colegio/server/lib/store"

	"github.com/pkg/errors"
)


func RefreshToken(ctx context.Context, refreshTokenString string) (*models.UserToken, []*httpresponses.ValidationError, error) {
	dbStore := store.NewStoreDefault()
	authorizer := selfauthorizer.NewSelfAuthorizer()

	return refreshToken(ctx, dbStore, authorizer, refreshTokenString)
}

func refreshToken(ctx context.Context, dbStore store.Store, authorizer auth.Auth, refreshTokenString string) (*models.UserToken, []*httpresponses.ValidationError, error) {
	parsedToken, err := authorizer.Validate(refreshTokenString)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	ctx, err = auth.DecorateContext(ctx, parsedToken)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	userDB, err := dbStore.GetUser(ctx, userID)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	user := userDB.ToModel()

	if !user.Active {
		return nil, nil, errors.New("usuario inactivo")
	}

	currentTokenVersion, err := auth.GetTokenVersion(ctx, parsedToken)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if currentTokenVersion != user.TokenVersion {
		return nil, nil, errors.New("token inválido")
	}

	accessToken, err := authorizer.GetSignedAccessToken(user.Role, user.Name, user.LastName, user.Email, user.ID.Hex())
	if err != nil {
		return nil, nil, err
	}

	newRefreshToken, err := authorizer.GetSignedRefreshToken(user.ID.Hex(), currentTokenVersion)
	if err != nil {
		return nil, nil, err
	}

	userToken := &models.UserToken{
		ID:           user.ID,
		Name:         user.Name,
		LastName:     user.LastName,
		Email:        user.Email,
		Rol:          user.Role,
		Token:        accessToken,
		RefreshToken: newRefreshToken,
		Avatar:       user.Avatar,
	}

	return userToken, nil, nil
}
