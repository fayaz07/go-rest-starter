package config

import (
	"flag"
	log "go-rest-starter/src/core/logger"
	"go-rest-starter/src/utils/constants"
	"go-rest-starter/src/utils/helpers"
	"sync"
)

const ALLOWED_ENVIRONMENTS_MSG = "Allowed values: dev, prod, test, staging"

var ALLOWED_ENVIRONMENTS = [...]string{constants.TEST_ENV, constants.DEV_ENV, constants.PROD_ENV, constants.STAGING_ENV}

var c_env string = constants.PROD_ENV
var _onceForEnv sync.Once

func getCurrentEnvironment() string {
	_onceForEnv.Do(func() {
		env_pointer := flag.String(constants.ENV_KEY, constants.PROD_ENV, ALLOWED_ENVIRONMENTS_MSG)
		flag.Parse()
		c_env = *env_pointer
	})
	return validateEnv(c_env)
}

func validateEnv(env string) string {
	if !helpers.Contains(ALLOWED_ENVIRONMENTS[:], env) {
		log.E("Environment not allowed")
		panic("Environment not allowed")
	}
	return string(env)
}
