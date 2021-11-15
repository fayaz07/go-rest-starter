package core

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func InitApplication() {
	appConfig := GetAppConfig()
	fmt.Println("#v", appConfig)

	GetDbConnection(appConfig.DB)

	InitRedis()

	server = gin.Default()
	InitializeServer(server)
}
