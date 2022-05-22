package routers

import (
	controllers "go-rest-starter/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

var appStatusController = controllers.GetAppStatusController()

func AddAppStatusRoutes(rg *gin.RouterGroup) {

	appStatus := rg.Group("/app")

	appStatus.GET("/hello", appStatusController.SayHello())

}
