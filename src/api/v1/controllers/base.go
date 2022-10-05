package controllers

import (
	wrappers "go-rest-starter/src/api/wrappers"
)

var resWrapper wrappers.ResponseWrapper = wrappers.GetResponseWrapper()

type BaseController struct{}
