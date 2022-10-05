package config

import (
	appTypes "go-rest-starter/src/core/types"
	"go-rest-starter/src/utils/constants"
	"os"
)

func loadSendGridConfig() *appTypes.SendGridConfig {
	return &appTypes.SendGridConfig{
		SgAPIKey:   os.Getenv(constants.SG_API_KEY),
		SgEmail:    os.Getenv(constants.SG_EMAIL),
		SgUsername: os.Getenv(constants.SG_USERNAME),
	}
}
