package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvFilePath(t *testing.T) {
	assert := assert.New(t)

	mockHome := "mock/home"

	t.Setenv("HOME", mockHome)

	assert.Equal(mockHome+"/dotenvs/"+APP_NAME+"/"+PROD_ENV+"/.env", getEnvFilePath())
}

func TestGetEnvFilePathWithInvalidHomeDir(t *testing.T) {
	assert := assert.New(t)

	mockHome := ""

	t.Setenv("HOME", mockHome)

	assert.Panics(func() { getEnvFilePath() })
}

func TestLoadEnvFileWithInvalidPath(t *testing.T) {
	assert := assert.New(t)

	mockHome := "mock/home"

	t.Setenv("HOME", mockHome)

	assert.Panics(func() { loadEnvFile() })
}

func TestLoadEnvFileWithValidPath(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() { loadEnvFile() })
}
