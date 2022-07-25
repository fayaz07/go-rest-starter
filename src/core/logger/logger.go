package logger

import (
	config "go-rest-starter/src/core/config"
	"os"

	"github.com/sirupsen/logrus"
)

var logger logrus.Logger

func SetupLogging() {
	logger = *logrus.New()
	setupOutput()
}

func setupOutput() {
	cEnv := config.GetAppConfig().AppEnv
	if cEnv == config.PROD_ENV {
		file, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.Out = file
		} else {
			panic("Failed to log to file, using default stderr")
		}
	} else {
		logger.Out = os.Stdout
	}
}

func I(d ...interface{}) {
	logger.Infoln(d)
}

func W(d ...interface{}) {
	logger.Warnln(d)
}

func E(d ...interface{}) {
	logger.Errorln(d)
}
