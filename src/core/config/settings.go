package config

import (
	appTypes "go-rest-starter/src/core/types"
	"log"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
)

var appSettings appTypes.AppSettings

const settingsFileName = "settings.properties"

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
	log.Println("App settings loaded from " + settingsFileName)
}

func GetAppSettings() appTypes.AppSettings {
	return appSettings
}
