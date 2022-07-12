package config

import (
	appTypes "go-rest-starter/src/core/types"
	"os"
)

func loadSendGridConfig() *appTypes.SendGridConfig {
	return &appTypes.SendGridConfig{
		SgAPIKey:   os.Getenv(sgAPIKey),
		SgEmail:    os.Getenv(sgEmail),
		SgUsername: os.Getenv(sgUsername),
	}
}