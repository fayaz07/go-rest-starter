package routers

import (
	controllers "go-rest-starter/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

// AddAppStatusRoutes will add routes regarding app's setup & health
func AddAppStatusRoutes(rg *gin.RouterGroup) {

	appStatus := rg.Group("/app")

	helloController := controllers.GetInstance()

	appStatus.GET("/hello", helloController.Hello())

}
