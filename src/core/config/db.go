package config

import (
	"fmt"
	"os"

	appTypes "go-rest-starter/src/core/types"
)

const DB_NAME_TEMPLATE = "%s-%s"

func loadDbConfig(appEnv string) *appTypes.DatabaseConfig {
	appDbName := os.Getenv(dBName)

	if appEnv != PROD_ENV {
		appDbName = fmt.Sprintf(DB_NAME_TEMPLATE, appDbName, appEnv)
	}

	return &appTypes.DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(dBHost),
		DbPort:     os.Getenv(dBPort),
		DbUsername: os.Getenv(dBUsername),
		DbPassword: os.Getenv(dBPassword),
	}
}
