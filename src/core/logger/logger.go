package logger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _loggerOnce sync.Once
var logger *zap.Logger = buildTestLogger()

func GetLogger() *zap.Logger {
	return logger
}

// --- info methods
func i(d string) {
	c := getCaller()
	go func() {
		logger.Info(d, callerField(c))
	}()
}

func I(d string) {
	i(d)
}

func If(template string, d ...interface{}) {
	i(fmt.Sprintf(template, d...))
}

// --- error methods
func e(d string) {
	c := getCaller()
	go func() {
		logger.Error(d, callerField(c))
	}()
}

func E(d string) {
	e(d)
}

func Ef(template string, d ...interface{}) {
	e(fmt.Sprintf(template, d...))
}

// --- warn methods
func w(d string) {
	c := getCaller()
	go func() {
		logger.Warn(d, callerField(c))
	}()
}

func W(d string) {
	w(d)
}

func Wf(template string, d ...interface{}) {
	w(fmt.Sprintf(template, d...))
}

// --- Fatal methods
func f(d string) {
	c := getCaller()
	go func() {
		logger.Fatal(d, callerField(c))
	}()
}

func F(d string) {
	f(d)
}

func Ff(template string, d ...interface{}) {
	f(fmt.Sprintf(template, d...))
}

// --- Panic methods
func p(d string) {
	c := getCaller()
	go func() {
		logger.Panic(d, callerField(c))
	}()
}

func P(d string) {
	p(d)
}

func Pf(template string, d ...interface{}) {
	p(fmt.Sprintf(template, d...))
}

func callerField(c string) zap.Field {
	return zapcore.Field{Key: "c", Type: zapcore.StringType, String: trimFilePath(c)}
}
