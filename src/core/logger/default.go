package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func buildTestLogger() *zap.Logger {
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	consoleDebugging := zapcore.Lock(os.Stdout)

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	return zap.New(zapcore.NewTee(zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority)))
}
