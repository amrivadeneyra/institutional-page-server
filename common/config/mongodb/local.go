package mongodbconfig

import "colegio/server/common/utils"

func local() *MongoDB {
	return &MongoDB{
		URL:      utils.GetEnvVar("DB_URL", false),
		UserName: utils.GetEnvVar("DB_USER", false),
		Password: utils.GetEnvVar("DB_PWD", false),
		DBName:   utils.GetEnvVar("DB_NAME", false),
		MaxPool:  utils.GetEnvVarInt("DB_MAX_POOL", false),
		UseSSL:   utils.GetEnvVarBool("DB_SSL", false),
		Timeout:  5,
	}
}
