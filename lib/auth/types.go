package auth

import "github.com/lestrrat-go/jwx/v2/jwt"

type Auth interface {
	Validate(token string) (jwt.Token, error)
	GetSignedAccessToken(rol, name, lastName, email, userID string) (string, error)
	GetSignedRefreshToken(userID string, tokenVersion int64) (string, error)
}
