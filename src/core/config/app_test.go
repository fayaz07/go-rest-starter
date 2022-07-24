package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAppConfigWithEmptyValue(t *testing.T) {
	t.Setenv(APP_PORT, "")

	assert := assert.New(t)

	env, port := loadAppConfig()

	assert.Equal(port, 7000)
	assert.Equal(env, PROD_ENV)
}

func TestLoadAppConfigWithCustomValues(t *testing.T) {
	t.Setenv(APP_PORT, "5000")

	assert := assert.New(t)

	env, port := loadAppConfig()

	assert.Equal(5000, port)
	assert.Equal(PROD_ENV, env)
}

func TestGetAppConfig(t *testing.T) {
	assert := assert.New(t)

	result := GetAppConfig()

	assert.Equal(true, result.Initiated)
	assert.Equal(PROD_ENV, result.AppEnv)
	assert.Equal(7000, result.AppPort)
}

func TestGetAppConfigShouldNotReInitialize(t *testing.T) {
	assert := assert.New(t)

	result := GetAppConfig()

	assert.Equal(true, result.Initiated)
	assert.Equal(PROD_ENV, result.AppEnv)
	assert.Equal(7000, result.AppPort)

	t.Setenv(APP_PORT, "3000")

	result2 := GetAppConfig()

	assert.NotEqual(3000, result2.AppPort)
}
