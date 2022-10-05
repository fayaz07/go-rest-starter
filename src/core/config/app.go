package config

import (
	"os"
	"strconv"

	log "go-rest-starter/src/core/logger"
	appTypes "go-rest-starter/src/core/types"
	"go-rest-starter/src/utils/constants"
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

	appEnv := getCurrentEnvironment()

	log.If("Running on %s environment", appEnv)

	loadSettings()
	log.SetupLogging(GetAppSettings())

	log.I("Initialising application configuration")

	loadEnvFile()

	appPort := loadAppConfig()

	_config = appTypes.AppConfig{
		Initiated: true,
		AppEnv:    appEnv,
		AppPort:   appPort,
		DB:        loadDbConfig(appEnv),
		SG:        loadSendGridConfig(),
		Jwt:       loadJwtConfig(),
	}
}

func loadAppConfig() int {
	appPort, err := strconv.Atoi(os.Getenv(constants.APP_PORT))
	if err != nil {
		log.If("App port can't be parsed, setting to default port: %d", constants.APP_DEFAULT_PORT)
		appPort = constants.APP_DEFAULT_PORT
	}
	return appPort
}
