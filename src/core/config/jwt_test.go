package config

import (
	"go-rest-starter/src/utils/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJwtConfigMinutes(t *testing.T) {
	assert := assert.New(t)

	mockAccessTokenExp := "10m"
	mockRefreshTokenExp := "20m"

	t.Setenv(constants.ACCESS_TOKEN_EXPIRY_KEY, mockAccessTokenExp)
	t.Setenv(constants.REFRESH_TOKEN_EXPIRY_KEY, mockRefreshTokenExp)

	result := loadJwtConfig()

	assert.Equal(MinutesToSeconds(10), result.AccessTokenExpiry)
	assert.Equal(MinutesToSeconds(20), result.RefreshTokenExpiry)
}

func TestLoadJwtConfigHours(t *testing.T) {
	assert := assert.New(t)

	mockAccessTokenExp := "1h"
	mockRefreshTokenExp := "2h"

	t.Setenv(constants.ACCESS_TOKEN_EXPIRY_KEY, mockAccessTokenExp)
	t.Setenv(constants.REFRESH_TOKEN_EXPIRY_KEY, mockRefreshTokenExp)

	result := loadJwtConfig()

	assert.Equal(HoursToSeconds(1), result.AccessTokenExpiry)
	assert.Equal(HoursToSeconds(2), result.RefreshTokenExpiry)
}

func TestLoadJwtConfigDefault(t *testing.T) {
	assert := assert.New(t)

	result := loadJwtConfig()

	assert.Equal(MinutesToSeconds(10), result.AccessTokenExpiry)
	assert.Equal(HoursToSeconds(2), result.RefreshTokenExpiry)
}
