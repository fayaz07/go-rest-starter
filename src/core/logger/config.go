package logger

import (
	"fmt"
	"go-rest-starter/src/core/types"
	"time"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const logFileFormat = "2006_January_02"

func SetupLogging(settings types.AppSettings, env string) {
	_loggerOnce.Do(func() {
		logger = *logrus.New()

		setupOutput(settings)
	})
}

func setupOutput(settings types.AppSettings) {
	logger.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", settings.LogsDir, time.Now().Format(logFileFormat)),
		MaxSize:    49,
		MaxBackups: 1,
		MaxAge:     1,
		Compress:   true,
	})

	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "t",
		logrus.FieldKeyLevel: "l",
		logrus.FieldKeyMsg:   "m",
		logrus.FieldKeyFunc:  "c",
	}

	logger.SetFormatter(&logrus.TextFormatter{QuoteEmptyFields: true, FullTimestamp: true, FieldMap: fieldMap})

	logger.Infoln("------------------------------------------------------------------------------------------------------------------")
	logger.Infof("--------------------------------------%s--------------------------------------------\n", time.Now().Format(logFileFormat))
	logger.Infoln("------------------------------------------------------------------------------------------------------------------")
}
