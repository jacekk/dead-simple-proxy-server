package worker

import (
	"fmt"
	"os"
	"time"
)

type logger interface {
	Info(string, ...interface{})
}

type loggerImpl struct{}

func (l *loggerImpl) Info(baseFormat string, args ...interface{}) {
	now := time.Now().Format("2006/01/02 - 15:04:05")
	format := fmt.Sprintf("%s %s | %s \n", loggerPrefix, now, baseFormat)
	fmt.Fprintf(os.Stdout, format, args...)
}
