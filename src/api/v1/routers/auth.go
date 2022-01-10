package routers

import (
	controllers "go-rest-starter/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

var authController = controllers.GetAuthController()

func AddAuthRoutes(rg *gin.RouterGroup) {

	authRoute := rg.Group("/auth")

	authRoute.GET("/init", authController.Init())

}
