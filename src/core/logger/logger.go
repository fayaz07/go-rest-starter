package logger

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var _loggerOnce sync.Once
var logger logrus.Logger

// --- info methods
func i(d interface{}) {
	c := getCaller()
	go func() {
		logger.WithField("c", trimFilePath(c)).Infoln(d)
	}()
}

func I(d interface{}) {
	i(d)
}

func If(template string, d ...interface{}) {
	i(fmt.Sprintf(template, d...))
}

// --- error methods
func e(d interface{}) {
	c := getCaller()
	go func() {
		logger.WithField("c", trimFilePath(c)).Errorln(d)
	}()
}

func E(d interface{}) {
	e(d)
}

func Ef(template string, d ...interface{}) {
	e(fmt.Sprintf(template, d...))
}

// --- warn methods
func w(d interface{}) {
	c := getCaller()
	go func() {
		logger.WithField("c", trimFilePath(c)).Warnln(d)
	}()
}

func W(d interface{}) {
	w(d)
}

func Wf(template string, d ...interface{}) {
	w(fmt.Sprintf(template, d...))
}

// --- Fatal methods
func f(d interface{}) {
	c := getCaller()
	go func() {
		logger.WithField("c", trimFilePath(c)).Fatalln(d)
	}()
}

func F(d interface{}) {
	f(d)
}

func Ff(template string, d ...interface{}) {
	f(fmt.Sprintf(template, d...))
}

// --- Panic methods
func p(d interface{}) {
	c := getCaller()
	go func() {
		logger.WithField("c", trimFilePath(c)).Panicln(d)
	}()
}

func P(d interface{}) {
	p(d)
}

func Pf(template string, d ...interface{}) {
	p(fmt.Sprintf(template, d...))
}
