package config

import (
	appTypes "go-rest-starter/src/core/types"
	"os"
)

func loadSendGridConfig() *appTypes.SendGridConfig {
	return &appTypes.SendGridConfig{
		SgAPIKey:   os.Getenv(SG_API_KEY),
		SgEmail:    os.Getenv(SG_EMAIL),
		SgUsername: os.Getenv(SG_USERNAME),
	}
}
