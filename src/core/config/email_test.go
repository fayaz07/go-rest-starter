package config

import (
	"go-rest-starter/src/utils/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSendGridConfig(t *testing.T) {
	assert := assert.New(t)

	mockApiKey := "API0de9b0cdf3921e218d985a10a85b64da"
	mockEmail := "mock@mock.com"
	mockUsername := "mock_username"

	t.Setenv(constants.SG_API_KEY, mockApiKey)
	t.Setenv(constants.SG_EMAIL, mockEmail)
	t.Setenv(constants.SG_USERNAME, mockUsername)

	result := loadSendGridConfig()
	assert.Equal(mockApiKey, result.SgAPIKey)
	assert.Equal(mockEmail, result.SgEmail)
	assert.Equal(mockUsername, result.SgUsername)
}
