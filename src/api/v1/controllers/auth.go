package controllers

import (
	"net/http"

	reqModels "go-rest-starter/src/api/models/req/auth"
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
func (auth) Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req reqModels.AuthInitReq

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, nil))
			return
		}

		helpers.ParseExtraClientInfo(c, &req)

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, resWrapper.Error(constants.Invalid_Request_Body, err))
			return
		}

		_authRepo.Init(req)

		c.JSON(http.StatusOK, resWrapper.Success("This will init auth", nil))
		c.Done()
	}
}
