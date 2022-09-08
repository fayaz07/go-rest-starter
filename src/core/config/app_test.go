package config

import (
	"go-rest-starter/src/utils/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAppConfigWithEmptyValue(t *testing.T) {
	assert := assert.New(t)

	port := loadAppConfig()

	assert.Equal(port, 7000)
}

func TestLoadAppConfigWithCustomValues(t *testing.T) {
	t.Setenv(constants.APP_PORT, "5000")

	assert := assert.New(t)

	port := loadAppConfig()

	assert.Equal(5000, port)
}

func TestGetAppConfig(t *testing.T) {
	assert := assert.New(t)

	result := GetAppConfig()

	assert.Equal(true, result.Initiated)
	assert.Equal(constants.PROD_ENV, result.AppEnv)
	assert.Equal(7000, result.AppPort)
}

func TestGetAppConfigShouldNotReInitialize(t *testing.T) {
	assert := assert.New(t)

	result := GetAppConfig()

	assert.Equal(true, result.Initiated)
	assert.Equal(constants.PROD_ENV, result.AppEnv)
	assert.Equal(7000, result.AppPort)

	t.Setenv(constants.APP_PORT, "3000")

	result2 := GetAppConfig()

	assert.NotEqual(3000, result2.AppPort)
}
