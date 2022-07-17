package config

import (
	appTypes "go-rest-starter/src/core/types"
	"os"
	"time"
)

const (
	ONE_MINUTE   = 60
	HALF_AN_HOUR = ONE_MINUTE * 30
	ONE_HOUR     = ONE_MINUTE * 60
)

func HoursToSeconds(hrs int) int {
	return ONE_HOUR * hrs
}

func MinutesToSeconds(mins int) int {
	return ONE_MINUTE * mins
}

func loadJwtConfig() *appTypes.JwtConfig {
	var a_exp int = HoursToSeconds(1)
	var b_exp int = HoursToSeconds(12)

	a, err := time.ParseDuration(os.Getenv(ACCESS_TOKEN_EXPIRY_KEY))

	if err == nil {
		a_exp = int(a.Seconds())
	}

	b, err := time.ParseDuration(os.Getenv(REFRESH_TOKEN_EXPIRY_KEY))

	if err == nil {
		b_exp = int(b.Seconds())
	}

	return &appTypes.JwtConfig{
		AccessTokenExpiry:  a_exp,
		RefreshTokenExpiry: b_exp,
	}
}
