package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Abstracting functions to be used only by those routers which require it
type auth BaseController

var _authInstance auth

func GetAuthController() auth {
	return _authInstance
}

// Controller functions implementation start here

func (auth) Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, resWrapper.Success("This will init auth", nil))
		c.Done()
	}
}
