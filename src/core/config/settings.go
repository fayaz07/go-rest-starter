package config

import (
	appTypes "go-rest-starter/src/core/types"
	"log"
	"path/filepath"

	"github.com/magiconair/properties"
)

var appSettings appTypes.AppSettings

const settingsFileName = "settings.properties"

func init() {
	loadSettings()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadSettings() {
	absolutePath, err := filepath.Abs(settingsFileName)
	check(err)

	p := properties.MustLoadFile(absolutePath, properties.UTF8)
	if err := p.Decode(&appSettings); err != nil {
		log.Panic(err)
	}
	log.Println("App settings loaded from " + settingsFileName)
}

func GetAppSettings() appTypes.AppSettings {
	return appSettings
}
