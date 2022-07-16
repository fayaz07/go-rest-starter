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

	appPort, err := strconv.Atoi(os.Getenv(appPort))
	if err != nil {
		log.Fatal("App port can't be parsed, setting to default port: 7000")
		appPort = 7000
	}

	log.Printf("Running on %s environment\n", appEnv)
	return appEnv, appPort
}
