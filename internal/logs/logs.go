package logs

import (
	"github.com/KIVUOS1999/easyLogs/pkg/models"
)

func Error(inp ...any) {
	log(models.Error, inp...)
}

func Info(inp ...any) {
	log(models.Info, inp...)
}

func Debug(inp ...any) {
	log(models.Debug, inp...)
}

func Warn(inp ...any) {
	log(models.Warn, inp...)
}

func Errorf(inp string, args ...any) {
	logf(models.Error, inp, args...)
}

func Infof(inp string, args ...any) {
	logf(models.Info, inp, args...)
}

func Debugf(inp string, args ...any) {
	logf(models.Debug, inp, args...)
}

func Warnf(inp string, args ...any) {
	logf(models.Warn, inp, args...)
}

func ErrorWithTrace(args ...any) {
	log(models.ErrorWithTrace, args...)
}
