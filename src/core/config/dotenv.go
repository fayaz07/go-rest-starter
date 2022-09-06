package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const ENV_FILE_PATH_STR = "%s/%s/%s/%s/%s"

func loadEnvFile() {
	log.Println("Loading .env file")
	filePath := getEnvFilePath()
	log.Printf("Env file path: %s\n", filePath)
	err := godotenv.Load(filePath)

	if err != nil {
		panic("Error loading .env file")
	}
}

func getEnvFilePath() string {
	appSettings := GetAppSettings()
	return fmt.Sprintf(
		ENV_FILE_PATH_STR,
		getHomeDir(),
		appSettings.DotEnvsDir,
		appSettings.AppName,
		getCurrentEnvironment(),
		appSettings.DotEnvFile,
	)
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dirname
}
