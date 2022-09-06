package config

import (
	log "go-rest-starter/src/core/logger"
	"os"

	"github.com/joho/godotenv"
)

func loadEnvFile() {
	log.I("Loading .env file")
	filePath := getEnvFilePath()
	log.If("Env file path: %s\n", filePath)
	err := godotenv.Load(filePath)

	if err != nil {
		log.P("Error loading .env file")
		panic("Error loading .env file")
	}
}

func getEnvFilePath() string {
	return GetAppSettings().DotEnvFile
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.P(err)
		panic(err)
	}
	return dirname
}
