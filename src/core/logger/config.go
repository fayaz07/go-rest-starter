package logger

import (
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogging() {
	_loggerOnce.Do(func() {
		logger = *logrus.New()

		setupOutput()
	})
}

func setupOutput() {
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "/Users/fayaz.mohammad/me/go-rest-starter/foo.log",
		MaxSize:    49,
		MaxBackups: 1,
		MaxAge:     7,
		Compress:   true,
	})

	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "t",
		logrus.FieldKeyLevel: "l",
		logrus.FieldKeyMsg:   "m",
		logrus.FieldKeyFunc:  "c",
	}

	logger.SetFormatter(&logrus.TextFormatter{QuoteEmptyFields: true, FullTimestamp: true, FieldMap: fieldMap})
}
