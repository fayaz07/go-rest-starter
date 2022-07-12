package config

import (
	"log"

	"github.com/joho/godotenv"
)

func loadEnvFile() {
	log.Println("Loading .env file")
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
}
