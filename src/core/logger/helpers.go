package logger

import (
	"fmt"
	"runtime"
	"strings"
)

func trimFilePath(c string) string {
	idx := strings.LastIndex(c, "src/")
	if idx != -1 {
		c = c[idx+4:]
	}
	return c
}

func getCaller() string {
	_, file, no, ok := runtime.Caller(3)
	if ok {
		return fmt.Sprintf("%s#%d", file, no)
	}
	return ""
}
