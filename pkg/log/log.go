package log

import (
	"github.com/KIVUOS1999/easyLogs/internal/constants"
	"github.com/KIVUOS1999/easyLogs/internal/logs"
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func Init(logLevel configs.LogLevel, style configs.LogFormat) {
	constants.LogConfig.LogLevel = logLevel
	constants.LogConfig.LogFormat = style
}

func init() {
	constants.LogConfig.LogLevel = configs.Debug
	constants.LogConfig.LogFormat = configs.ColoredLogs
}

func Error(inp ...any) {
	logs.Error(inp...)
}

func Info(inp ...any) {
	logs.Info(inp...)
}

func Debug(inp ...any) {
	logs.Debug(inp...)
}

func Warn(inp ...any) {
	logs.Warn(inp...)
}

func Errorf(args ...any) {
	logs.Errorf(args...)
}

func Infof(args ...any) {
	logs.Infof(args...)
}

func Debugf(args ...any) {
	logs.Debugf(args...)
}

func Warnf(args ...any) {
	logs.Warnf(args...)
}

func ErrorWithTrace(args ...any) {
	logs.ErrorWithTrace(args...)
}
