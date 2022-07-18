package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentEnvWithNoFlag(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(PROD_ENV, getCurrentEnvironment())
}

func TestValidateEnvWithInvalidFlag(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValue("Environment not allowed", func() { validateEnv("stage") })
}

func TestValidateEnvWithDevEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(DEV_ENV, validateEnv(DEV_ENV))
	assert.NotPanics(func() { validateEnv(DEV_ENV) })
}

func TestValidateEnvWithProdEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(PROD_ENV, validateEnv(PROD_ENV))
	assert.NotPanics(func() { validateEnv(PROD_ENV) })
}

func TestValidateEnvWithTestEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(TEST_ENV, validateEnv(TEST_ENV))
	assert.NotPanics(func() { validateEnv(TEST_ENV) })
}

func TestValidateEnvWithStagingEnv(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(STAGING_ENV, validateEnv(STAGING_ENV))
	assert.NotPanics(func() { validateEnv(STAGING_ENV) })
}
