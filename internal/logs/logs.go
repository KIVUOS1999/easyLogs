package logs

import (
	constants "github.com/KIVUOS1999/easyLogs/internal"
)

func Error(inp ...any) {
	log(constants.Error, inp...)
}

func Info(inp ...any) {
	log(constants.Info, inp...)
}

func Debug(inp ...any) {
	log(constants.Debug, inp...)
}

func Warn(inp ...any) {
	log(constants.Warn, inp...)
}

func Errorf(inp string, args ...any) {
	logf(constants.Error, inp, args...)
}

func Infof(inp string, args ...any) {
	logf(constants.Info, inp, args...)
}

func Debugf(inp string, args ...any) {
	logf(constants.Debug, inp, args...)
}

func Warnf(inp string, args ...any) {
	logf(constants.Warn, inp, args...)
}

func ErrorWithTrace(args ...any) {
	log(constants.ErrorWithTrace, args...)
}
