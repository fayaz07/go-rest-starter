package config

import (
	"go-rest-starter/src/utils/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentEnvWithNoFlag(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(constants.PROD_ENV, getCurrentEnvironment())
}

func TestValidateEnvWithInvalidFlag(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValue("Environment not allowed", func() { validateEnv("stage") })
}

func TestValidateEnvWithDevEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(constants.DEV_ENV, validateEnv(constants.DEV_ENV))
	assert.NotPanics(func() { validateEnv(constants.DEV_ENV) })
}

func TestValidateEnvWithProdEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(constants.PROD_ENV, validateEnv(constants.PROD_ENV))
	assert.NotPanics(func() { validateEnv(constants.PROD_ENV) })
}

func TestValidateEnvWithTestEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(constants.TEST_ENV, validateEnv(constants.TEST_ENV))
	assert.NotPanics(func() { validateEnv(constants.TEST_ENV) })
}

func TestValidateEnvWithStagingEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(constants.STAGING_ENV, validateEnv(constants.STAGING_ENV))
	assert.NotPanics(func() { validateEnv(constants.STAGING_ENV) })
}
