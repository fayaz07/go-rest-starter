package config

import (
	"os"

	appTypes "go-rest-starter/src/core/types"
)

func loadDbConfig(appEnv string) *appTypes.DatabaseConfig {
	appDbName := os.Getenv(dBName)

	if appEnv != PROD_ENV {
		appDbName = appDbName + "-" + appEnv
	}

	return &appTypes.DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(dBHost),
		DbPort:     os.Getenv(dBPort),
		DbUsername: os.Getenv(dBUsername),
		DbPassword: os.Getenv(dBPassword),
	}
}
