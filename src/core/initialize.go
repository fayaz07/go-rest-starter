package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func InitApplication() {
	appConfig := GetAppConfig()
	log.Println("#v", appConfig)

	log.Println("Connecting to database ....")
	InitializeDatabseConn(appConfig.DB)
	db := GetDbConnection()
	if db != nil {
		log.Println("Connected to database")
	} else {
		log.Println("Failed to connect to database")
	}

	InitRedis()

	server = gin.Default()
	log.Println("Starting server")
	InitializeServer(server)
	log.Println("Server started")
}
