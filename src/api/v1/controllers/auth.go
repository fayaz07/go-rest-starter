package controllers

import (
	"net/http"

	reqModels "go-rest-starter/src/api/models/req/auth"
	"go-rest-starter/src/utils/constants"

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
		var req reqModels.AuthInitReq

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, nil))
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, err))
			return
		}

		c.JSON(http.StatusOK, resWrapper.Success("This will init auth", nil))
		c.Done()
	}
}
