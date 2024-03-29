package config

import (
	"go-rest-starter/src/utils/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvFilePath(t *testing.T) {
	assert := assert.New(t)

	appSettings := GetAppSettings()

	assert.Contains(appSettings.DotEnvFile, "dotenvs/"+appSettings.AppName+"/"+constants.PROD_ENV+"/.env")
}

func TestLoadEnvFileWithValidPath(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() { loadEnvFile() })
}
