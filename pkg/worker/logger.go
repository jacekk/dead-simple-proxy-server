package worker

import (
	"fmt"
	"os"
)

type logger interface {
	Info(string, ...interface{})
}

type loggerImpl struct{}

func (l *loggerImpl) Info(baseFormat string, args ...interface{}) {
	format := fmt.Sprintf("%s %s \n", loggerPrefix, baseFormat)
	fmt.Fprintf(os.Stdout, format, args...)
}
