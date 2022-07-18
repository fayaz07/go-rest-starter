package config

import (
	"flag"
	"go-rest-starter/src/utils/helpers"
	"sync"
)

const ALLOWED_ENVIRONMENTS_MSG = "Allowed values: dev, prod, test, staging"

var ALLOWED_ENVIRONMENTS = [...]string{TEST_ENV, DEV_ENV, PROD_ENV, STAGING_ENV}

var c_env string = PROD_ENV
var _onceForEnv sync.Once

func getCurrentEnvironment() string {
	_onceForEnv.Do(func() {
		env_pointer := flag.String(envConst, PROD_ENV, ALLOWED_ENVIRONMENTS_MSG)
		flag.Parse()
		c_env = *env_pointer
	})
	return validateEnv(c_env)
}

func validateEnv(env string) string {
	if !helpers.Contains(ALLOWED_ENVIRONMENTS[:], env) {
		panic("Environment not allowed")
	}
	return string(env)
}
