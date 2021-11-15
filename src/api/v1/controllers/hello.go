package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	wrappers "go-rest-starter/src/api/wrappers"
)

var resWrapper wrappers.ResponseWrapper = wrappers.GetResponseWrapper()

// Abstracting functions to be used only by those routers which require it
type hello BaseController

var instance hello

func GetHelloController() hello {
	return instance
}

// Controller functions implementation start here

func (hello) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, resWrapper.Success("Hello", nil))
		c.Done()
	}
}
