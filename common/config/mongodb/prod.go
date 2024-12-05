package mongodbconfig

import "colegio/server/common/utils"

func prod() *MongoDB {
	return &MongoDB{
		URL:      utils.GetEnvVar("DB_URL", true),
		UserName: utils.GetEnvVar("DB_USER", true),
		Password: utils.GetEnvVar("DB_PWD", true),
		DBName:   utils.GetEnvVar("DB_NAME", true),
		MaxPool:  utils.GetEnvVarInt("DB_MAX_POOL", true),
		UseSSL:   utils.GetEnvVarBool("DB_SSL", true),
		Timeout:  5,
	}
}
