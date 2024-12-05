package config

import (
	mongodbconfig "colegio/server/common/config/mongodb"
	selfauthconfig "colegio/server/common/config/selfauth"
	"colegio/server/common/utils"

	"github.com/pkg/errors"
)

type RunConfig struct {
	SelfAuth *selfauthconfig.SelfAuth
	MongoDB  *mongodbconfig.MongoDB
}

func GetConfigDefault() *RunConfig {
	return GetConfig(utils.GetStage())
}

func GetConfig(stage utils.Stage) *RunConfig {
	return &RunConfig{
		SelfAuth: selfauthconfig.GetConfig(stage),
		MongoDB:  mongodbconfig.GetConfig(stage),
	}
}

func SwitchOnStage[T any](stage utils.Stage,
	prodValueFunc func() T,
	localValueFunc func() T,
) T {
	switch stage {
	case utils.Prod:
		return prodValueFunc()
	case utils.Local:
		return localValueFunc()
	default:
		panic(errors.Errorf("Not supported stage: %v", stage))
	}
}
