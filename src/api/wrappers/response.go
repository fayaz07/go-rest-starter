package wrappers

import (
	consts "go-rest-starter/src/utils/constants"

	"github.com/gin-gonic/gin"
)

type ResponseWrapper BaseWrapper

var instance ResponseWrapper

func GetResponseWrapper() ResponseWrapper {
	return instance
}

func (ResponseWrapper) Success(message string, data interface{}) gin.H {
	return handledSuccessResponse(message, data)
}

func handledSuccessResponse(message string, data interface{}) gin.H {
	if data == nil {
		return gin.H{consts.Status: consts.Success, consts.Message: message}
	}
	return gin.H{consts.Status: consts.Success, consts.Message: message, consts.Data: data}
}

func (ResponseWrapper) Error(message string, error interface{}) gin.H {
	return handledErrorResponse(message, error)
}

func handledErrorResponse(message string, error interface{}) gin.H {
	if error == nil {
		return gin.H{consts.Status: consts.Failed, consts.Message: message}
	}
	return gin.H{consts.Status: consts.Failed, consts.Message: message, consts.Error: error}
}
