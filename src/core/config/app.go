package config

import (
	"log"
	"os"
	"strconv"

	appTypes "go-rest-starter/src/core/types"
)

var _config appTypes.AppConfig

func GetAppConfig() *appTypes.AppConfig {
	if _config.Initiated {
		return &_config
	}
	initAppConfig()
	return &_config
}

func initAppConfig() {
	log.Println("Initialising application configuration")

	loadEnvFile()

	appEnv, appPort := loadAppConfig()

	_config = appTypes.AppConfig{
		Initiated: true,
		AppEnv:    appEnv,
		AppPort:   appPort,
		DB:        loadDbConfig(appEnv),
		SG:        loadSendGridConfig(),
		Jwt:       loadJwtConfig(),
	}
}

func loadAppConfig() (string, int) {
	appEnv := getCurrentEnvironment()

	appPort, err := strconv.Atoi(os.Getenv(APP_PORT))
	if err != nil {
		log.Printf("App port can't be parsed, setting to default port: %d\n", APP_DEFAULT_PORT)
		appPort = APP_DEFAULT_PORT
	}

	log.Printf("Running on %s environment\n", appEnv)
	return appEnv, appPort
}
