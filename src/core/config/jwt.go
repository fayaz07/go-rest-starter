package config

import (
	appTypes "go-rest-starter/src/core/types"
	"os"
	"strconv"
)

func loadJwtConfig() *appTypes.JwtConfig {
	a, err := strconv.Atoi(os.Getenv(accessTokenExpiry))
	if err != nil {
		a = 1 * 60 * 60
	}

	b, err := strconv.Atoi(os.Getenv(refreshTokenExpiry))
	if err != nil {
		b = 12 * 60 * 60
	}

	return &appTypes.JwtConfig{
		AccessTokenExpiry:  a,
		RefreshTokenExpiry: b,
	}
}
