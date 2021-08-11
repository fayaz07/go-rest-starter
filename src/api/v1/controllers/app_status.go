package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Abstracting functions to be used only by those routers which require it
type hello BaseController

var instance hello

func GetInstance() hello {
	return instance
}

// Controller functions implementation start here

// Hello this will
func (hello) Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "Hello"})
		c.Done()
	}
}
