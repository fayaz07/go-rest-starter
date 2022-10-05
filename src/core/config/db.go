package config

import (
	"fmt"
	"os"

	appTypes "go-rest-starter/src/core/types"
	"go-rest-starter/src/utils/constants"
)

const DB_NAME_TEMPLATE = "%s-%s"

func loadDbConfig(appEnv string) *appTypes.DatabaseConfig {
	appDbName := os.Getenv(constants.DB_NAME)

	if appEnv != constants.PROD_ENV {
		appDbName = fmt.Sprintf(DB_NAME_TEMPLATE, appDbName, appEnv)
	}

	return &appTypes.DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(constants.DB_HOST),
		DbPort:     os.Getenv(constants.DB_PORT),
		DbUsername: os.Getenv(constants.DB_USERNAME),
		DbPassword: os.Getenv(constants.DB_PASSWORD),
	}
}
