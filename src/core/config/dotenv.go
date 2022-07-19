package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const APP_NAME = "go-rest-starter"
const ENV_FILE_NAME = ".env"
const DOT_ENVS_FOLDER = "dotenvs"
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
	return fmt.Sprintf(ENV_FILE_PATH_STR, getHomeDir(), DOT_ENVS_FOLDER, APP_NAME, getCurrentEnvironment(), ENV_FILE_NAME)
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dirname
}
