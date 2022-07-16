package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSendGridConfig(t *testing.T) {
	assert := assert.New(t)

	mockApiKey := "API0de9b0cdf3921e218d985a10a85b64da"
	mockEmail := "mock@mock.com"
	mockUsername := "mock_username"

	t.Setenv(SgAPIKey, mockApiKey)
	t.Setenv(SgEmail, mockEmail)
	t.Setenv(SgUsername, mockUsername)

	result := loadSendGridConfig()
	assert.Equal(result.SgAPIKey, mockApiKey)
	assert.Equal(result.SgEmail, mockEmail)
	assert.Equal(result.SgUsername, mockUsername)
}
