package core

import (
	log "go-rest-starter/src/core/logger"

	"github.com/gin-gonic/gin"

	config "go-rest-starter/src/core/config"
	helpers "go-rest-starter/src/utils/helpers"
)

var server *gin.Engine

func InitApplication() {
	appConfig := config.GetAppConfig()
	log.I(helpers.Pretty(appConfig))

	InitRedis()

	log.I("Connecting to database ....")
	InitializeDatabseConn(appConfig.DB)
	db := GetDbConnection()
	if db != nil {
		log.I("Connected to database")
	} else {
		log.F("Failed to connect to database")
	}

	server = gin.Default()
	log.I("Starting server")
	InitializeServer(server)
	log.I("Server started")
}
