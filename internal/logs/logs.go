package logs

import (
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func Error(inp ...any) {
	formatDecider(configs.Error, false, inp...)
}

func Info(inp ...any) {
	formatDecider(configs.Info, false, inp...)
}

func Debug(inp ...any) {
	formatDecider(configs.Debug, false, inp...)
}

func Warn(inp ...any) {
	formatDecider(configs.Warn, false, inp...)
}

func Errorf(args ...any) {
	formatDecider(configs.Error, true, args...)
}

func Infof(args ...any) {
	formatDecider(configs.Info, true, args...)
}

func Debugf(args ...any) {
	formatDecider(configs.Debug, true, args...)
}

func Warnf(args ...any) {
	formatDecider(configs.Warn, true, args...)
}

func ErrorWithTrace(args ...any) {
	formatDecider(configs.ErrorWithTrace, false, args...)
}
