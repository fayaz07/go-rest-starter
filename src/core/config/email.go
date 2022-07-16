package config

import (
	appTypes "go-rest-starter/src/core/types"
	"os"
)

func loadSendGridConfig() *appTypes.SendGridConfig {
	return &appTypes.SendGridConfig{
		SgAPIKey:   os.Getenv(SgAPIKey),
		SgEmail:    os.Getenv(SgEmail),
		SgUsername: os.Getenv(SgUsername),
	}
}
