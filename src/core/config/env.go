package config

import (
	"flag"
	helpers "go-rest-starter/src/utils/helpers"
	"sync"
)

const ALLOWED_ENVIRONMENTS_MSG = "Allowed values: dev, prod, test, staging"

var ALLOWED_ENVIRONMENTS = [...]string{TEST_ENV, DEV_ENV, PROD_ENV, STAGING_ENV}

var c_env string = PROD_ENV
var _envInitialized bool = false
var _onceForEnv sync.Once

func getCurrentEnvironment() string {
	if _envInitialized {
		return c_env
	}
	_onceForEnv.Do(func() {
		env_pointer := flag.String(envConst, PROD_ENV, ALLOWED_ENVIRONMENTS_MSG)
		_envInitialized = true
		flag.Parse()
		c_env = *env_pointer

		if !helpers.Contains(ALLOWED_ENVIRONMENTS[:], c_env) {
			panic("Environment not allowed, " + c_env)
		}
	})
	return string(c_env)
}
