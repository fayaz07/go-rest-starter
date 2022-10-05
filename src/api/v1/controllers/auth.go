package controllers

import (
	"net/http"

	reqModels "go-rest-starter/src/api/models/req"
	"go-rest-starter/src/api/repository"
	"go-rest-starter/src/utils/constants"

	helpers "go-rest-starter/src/utils/helpers"

	"github.com/gin-gonic/gin"
)

// Abstracting functions to be used only by those routers which require it
type auth BaseController

var _authInstance auth

func GetAuthController() auth {
	return _authInstance
}

var _authRepo = repository.GetAuthRepo()

// Controller functions implementation start here
// TODO: Add logged in check
func (auth) InitSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _authRepo.AuthExchangeTokenExists(c.ClientIP()) {
			c.JSON(http.StatusAlreadyReported, resWrapper.Error("Session is already active", nil))
			return
		}
		var req reqModels.AuthInitReq

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, err.Error()))
			return
		}

		helpers.ParseExtraClientInfo(c, &req)

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, err))
			return
		}

		result, err := _authRepo.InitSession(req)
		if err != nil {
			c.JSON(http.StatusForbidden, resWrapper.Error("Session init failed", err.Error()))
		} else {
			c.JSON(http.StatusOK, resWrapper.Success("Session initialized", gin.H{"token": result}))
		}

		c.Done()
	}
}
