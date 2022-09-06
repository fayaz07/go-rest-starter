package logger

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var _loggerOnce sync.Once
var logger logrus.Logger

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
}

func I(d interface{}) {
	go func() {
		logger.Infoln(d)
	}()
}

func If(template string, d ...interface{}) {
	go func() {
		I(fmt.Sprintf(template, d...))
	}()
}

func W(d interface{}) {
	go func() {
		logger.Warnln(d)
	}()
}

func Wf(template string, d ...interface{}) {
	go func() {
		W(fmt.Sprintf(template, d...))
	}()
}

func E(d interface{}) {
	go func() {
		logger.Errorln(d)
	}()
}

func Ef(template string, d ...interface{}) {
	go func() {
		E(fmt.Sprintf(template, d...))
	}()
}
