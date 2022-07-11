package config

import (
	"log"
	"os"

	appTypes "go-rest-starter/src/core/types"
)

func loadDbConfig(appEnv string) *appTypes.DatabaseConfig {
	// get the database name from .env
	appDbName := os.Getenv(dBName)

	// structure port, database name according to environment
	if appEnv == "prod" {
	} else if appEnv == "test" {
		appDbName = appDbName + "-test"
	} else if appEnv == "dev" {
		appDbName = appDbName + "-dev"
	} else if appEnv == "staging" {
		appDbName = appDbName + "-staging"
	}

	log.Println(appDbName)

	return &appTypes.DatabaseConfig{
		DbName:     appDbName,
		DbHost:     os.Getenv(dBHost),
		DbPort:     os.Getenv(dBPort),
		DbUsername: os.Getenv(dBUsername),
		DbPassword: os.Getenv(dBPassword),
	}
}
