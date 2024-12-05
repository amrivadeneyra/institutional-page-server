package selfauthorizer

import (
	"colegio/server/common/config"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type SelfAuthorizer struct {
}

func NewSelfAuthorizer() *SelfAuthorizer {
	return &SelfAuthorizer{}
}

func (s *SelfAuthorizer) Validate(token string) (jwt.Token, error) {
	selfAuthConfig := config.GetConfigDefault().SelfAuth

	privkey, err := jwk.ParseKey([]byte(selfAuthConfig.AccessSecretKey))
	if err != nil {
		return nil, err
	}
	pubkey, err := jwk.PublicKeyOf(privkey)
	if err != nil {
		return nil, err
	}

	return jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, pubkey))
}

func (s *SelfAuthorizer) GetSignedAccessToken(role, name, lastName, email, userID string) (string, error) {
	selfAuthConfig := config.GetConfigDefault().SelfAuth

	token, err := jwt.NewBuilder().
		Issuer("colegio.com").
		IssuedAt(time.Now()).
		Claim("name", name).
		Claim("last_name", lastName).
		Claim("user_id", userID).
		Claim("email", email).
		Claim("role", role).
		Expiration(time.Now().Add(selfAuthConfig.AccessTokenExpiration)).
		Build()

	if err != nil {
		return "", err
	}

	privkey, err := jwk.ParseKey([]byte(selfAuthConfig.AccessSecretKey))
	if err != nil {
		return "", err
	}
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, privkey))
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

func (s *SelfAuthorizer) GetSignedRefreshToken(userID string, tokenVersion int64) (string, error) {
	selfAuthConfig := config.GetConfigDefault().SelfAuth

	token, err := jwt.NewBuilder().
		Issuer("colegio.com").
		IssuedAt(time.Now()).
		Claim("user_id", userID).
		Claim("token_version", tokenVersion).
		Expiration(time.Now().Add(selfAuthConfig.RefreshTokenExpiration)).
		Build()

	if err != nil {
		return "", err
	}

	privkey, err := jwk.ParseKey([]byte(selfAuthConfig.RefreshSecretKey))
	if err != nil {
		return "", err
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, privkey))
	if err != nil {
		return "", err
	}

	return string(signed), nil
}
