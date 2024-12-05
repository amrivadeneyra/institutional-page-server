package selfauthconfig

import (
	configutils "colegio/server/common/config/utils"
	"colegio/server/common/utils"
	"time"
)

type SelfAuth struct {
	AccessSecretKey        string
	AccessTokenExpiration  time.Duration
	RefreshSecretKey       string
	RefreshTokenExpiration time.Duration
}

func GetConfig(stage utils.Stage) *SelfAuth {
	return configutils.SwitchOnStage(stage,
		prod,
		local,
	)
}

func GetDefaultConfig() *SelfAuth {
	return GetConfig(utils.GetStage())
}
