package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSettings(t *testing.T) {
	assert := assert.New(t)

	loadSettings()

	appSettings := GetAppSettings()
	assert.NotNil(appSettings)
}
