package mongodbconfig

import (
	"colegio/server/common/utils"
	configutils "colegio/server/common/config/utils"
)

type MongoDB struct {
	URL      string
	UserName string
	Password string
	DBName   string
	MaxPool  int
	UseSSL   bool
	Timeout  int64
}

func GetConfig(stage utils.Stage) *MongoDB {
	return configutils.SwitchOnStage(stage,
		prod,
		local,
	)
}

func GetDefaultConfig() *MongoDB {
	return GetConfig(utils.GetStage())
}
