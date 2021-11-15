package routers

import (
	controllers "go-rest-starter/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

var helloController = controllers.GetHelloController()

func AddAppStatusRoutes(rg *gin.RouterGroup) {

	appStatus := rg.Group("/app")

	appStatus.GET("/hello", helloController.Hello())

}
