package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	wrappers "go-rest-starter/src/api/wrappers"
)

// Abstracting functions to be used only by those routers which require it
type hello BaseController

var instance hello

func GetHelloController() hello {
	return instance
}

var resWrapper wrappers.ResponseWrapper = wrappers.GetResponseWrapper()

// Controller functions implementation start here

// Hello this will return hello message
func (hello) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, resWrapper.Success("Hey buddy", nil))
		c.Done()
	}
}
