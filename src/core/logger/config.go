package logger

import (
	"fmt"
	"go-rest-starter/src/core/types"
	"go-rest-starter/src/utils/constants"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const logFileTemplate = "%s/%s.log"

func SetupLogging(settings types.AppSettings) {
	_loggerOnce.Do(func() {
		logger = zap.New(getZapConfig(settings))
		defer logger.Sync()
	})
}

func getZapConfig(settings types.AppSettings) zapcore.Core {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	logRotateWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename: getLogFileDir(settings),
		// so that log files are backed up after size limit is reached and are not deleted.
		MaxSize:    50, // megabytes
		MaxBackups: 0,
		MaxAge:     0,
	})

	return zapcore.NewTee(
		zapcore.NewCore(fileEncoder, logRotateWriter, zap.DebugLevel),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
	)
}

func getLogFileDir(settings types.AppSettings) string {
	if settings.LogsDir == "" {
		return "./temp/temp.log"
	}
	return fmt.Sprintf(logFileTemplate, settings.LogsDir, time.Now().Format(constants.LogFilenameFormat))
}

func logSetupDone() {

}
