package selfauthconfig

import (
	"colegio/server/common/utils"
	"time"
)

func prod() *SelfAuth {
	return &SelfAuth{
		AccessSecretKey:        utils.GetEnvVar("SECRET_KEY", true),
		AccessTokenExpiration:  time.Hour,
		RefreshSecretKey:       utils.GetEnvVar("REFRESH_SECRET_KEY", true),
		RefreshTokenExpiration: time.Hour * 168,
	}
}
