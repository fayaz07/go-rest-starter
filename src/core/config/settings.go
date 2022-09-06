package config

import (
	"fmt"
	appTypes "go-rest-starter/src/core/types"
	"log"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
)

var appSettings appTypes.AppSettings

const settingsFileName = "settings.properties"
const FOUR_DIR_APPEND_TEMPLATE = "%s/%s/%s/%s"
const THREE_DIR_APPEND_TEMPLATE = "%s/%s/%s"
const TWO_DIR_APPEND_TEMPLATE = "%s/%s"

func getSettingsFilePath() string {
	absolutePath, _ := filepath.Abs(settingsFileName)

	// because when this file is being called from Test, the file should be accessible
	return strings.Replace(absolutePath, "src/core/config/", "", 1)
}

func loadSettings() {
	filePath := getSettingsFilePath()

	p := properties.MustLoadFile(filePath, properties.UTF8)
	if err := p.Decode(&appSettings); err != nil {
		log.Panic(err)
	}

	homeDir := getHomeDir()
	currentEnv := getCurrentEnvironment()

	appSettings.DotEnvsDir = fmt.Sprintf(FOUR_DIR_APPEND_TEMPLATE, homeDir, appSettings.DotEnvsDir, appSettings.AppName, currentEnv)

	appSettings.DotEnvFile = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.DotEnvsDir, appSettings.DotEnvFile)

	appSettings.PublicDir = fmt.Sprintf(THREE_DIR_APPEND_TEMPLATE, homeDir, appSettings.PublicDir, appSettings.AppName)

	appSettings.UploadsPublicDir = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.PublicDir, appSettings.UploadsPublicDir)
	appSettings.StaticPublicDir = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.PublicDir, appSettings.StaticPublicDir)

	appSettings.LogsDir = fmt.Sprintf(FOUR_DIR_APPEND_TEMPLATE, homeDir, appSettings.LogsDir, appSettings.AppName, currentEnv)

	appSettings.KeysDir = fmt.Sprintf(FOUR_DIR_APPEND_TEMPLATE, homeDir, appSettings.KeysDir, appSettings.AppName, currentEnv)

	appSettings.AccessTokenPvtKey = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.KeysDir, appSettings.AccessTokenPvtKey)
	appSettings.AccessTokenPubKey = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.KeysDir, appSettings.AccessTokenPubKey)

	appSettings.RefreshTokenPvtKey = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.KeysDir, appSettings.RefreshTokenPvtKey)
	appSettings.RefreshTokenPubKey = fmt.Sprintf(TWO_DIR_APPEND_TEMPLATE, appSettings.KeysDir, appSettings.RefreshTokenPubKey)

	log.Println("App settings loaded from " + settingsFileName)
}

func GetAppSettings() appTypes.AppSettings {
	if appSettings.AppName != "go-rest-starter" {
		loadSettings()
	}

	return appSettings
}
