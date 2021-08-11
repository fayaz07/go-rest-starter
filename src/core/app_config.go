package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var _config AppConfig

func GetAppConfig() *AppConfig {
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
	_config = AppConfig{
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
	}
	return appEnv, appPort
}

func loadDbConfig(appEnv string) *DatabaseConfig {
	// get the database name from .env
	appDbName := os.Getenv(dBName)

	// structure port, database name according to environment
	if appEnv == "prod" {
	} else if appEnv == "test" {
		appDbName = appDbName + "-test"
	} else if appEnv == "dev" {
		appDbName = appDbName + "-dev"
	}

	return &DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(dBHost),
		DbPort:     os.Getenv(dBPort),
		DbUsername: os.Getenv(dBUsername),
		DbPassword: os.Getenv(dBPassword),
	}
}

func loadSendGridConfig() *SendGridConfig {
	return &SendGridConfig{
		SgAPIKey:   os.Getenv(sgAPIKey),
		SgEmail:    os.Getenv(sgEmail),
		SgUsername: os.Getenv(sgUsername),
	}
}

func loadJwtConfig() *JwtConfig {
	a, err := strconv.Atoi(os.Getenv(accessTokenExpiry))
	if err != nil {
		a = 1 * 60 * 60
	}

	b, err := strconv.Atoi(os.Getenv(refreshTokenExpiry))
	if err != nil {
		b = 12 * 60 * 60
	}

	return &JwtConfig{
		AccessTokenExpiry:  a,
		RefreshTokenExpiry: b,
	}
}
