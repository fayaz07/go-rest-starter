package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Abstracting functions to be used only by those routers which require it
type appStatus BaseController

var instance appStatus

func GetAppStatusController() appStatus {
	return instance
}

// Controller functions implementation start here
func (appStatus) SayHello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, resWrapper.Success("Hello", nil))
		c.Done()
	}
}
