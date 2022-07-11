package config

import (
	"log"
	"os"
	"strconv"

	appTypes "go-rest-starter/src/core/types"

	"github.com/joho/godotenv"
)

var _config appTypes.AppConfig

func GetAppConfig() *appTypes.AppConfig {
	if _config.Initiated {
		return &_config
	}
	initAppConfig()
	return &_config
}

// initAppConfig - pull data from .env
func initAppConfig() {
	log.Println("Initialising application configuration")

	loadEnvFile()

	appEnv, appPort := loadAppConfig()

	// initialize app config variable
	_config = appTypes.AppConfig{
		Initiated: true,
		AppEnv:    appEnv,
		AppPort:   appPort,
		DB:        loadDbConfig(appEnv),
		SG:        loadSendGridConfig(),
		Jwt:       loadJwtConfig(),
	}
}

func loadEnvFile() {
	log.Println("Loading .env file")
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}

func loadAppConfig() (string, int) {
	// get the environment from .env
	appEnv := os.Getenv(appEnv)

	// get the port from .env
	appPort, err := strconv.Atoi(os.Getenv(appPort))
	if err != nil {
		log.Fatal("App port can't be parsed, setting to default port: 7000")
		appPort = 7000
	}

	// structure port, database name according to environment
	if appEnv == "prod" {
		log.Print("Running on prod environment")
	} else if appEnv == "test" {
		log.Print("Running on test environment")
		appPort = appPort + 2
	} else if appEnv == "dev" {
		log.Print("Running on dev environment")
		appPort = appPort + 1
	} else if appEnv == "staging" {
		log.Print("Running on staging environment")
		appPort = appPort + 3
	}
	return appEnv, appPort
}
