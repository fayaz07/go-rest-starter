package config

import (
	"fmt"
	"os"

	appTypes "go-rest-starter/src/core/types"
)

const DB_NAME_TEMPLATE = "%s-%s"

func loadDbConfig(appEnv string) *appTypes.DatabaseConfig {
	appDbName := os.Getenv(DB_NAME)

	if appEnv != PROD_ENV {
		appDbName = fmt.Sprintf(DB_NAME_TEMPLATE, appDbName, appEnv)
	}

	return &appTypes.DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(DB_HOST),
		DbPort:     os.Getenv(DB_PORT),
		DbUsername: os.Getenv(DB_USERNAME),
		DbPassword: os.Getenv(DB_PASSWORD),
	}
}
