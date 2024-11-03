package log

import (
	constants "github.com/KIVUOS1999/easyLogs/internal"
	"github.com/KIVUOS1999/easyLogs/internal/logs"
	"github.com/KIVUOS1999/easyLogs/pkg/models"
)

func Init(logLevel models.LogLevel, style models.LogFormat) {
	constants.LogConfig.LogLevel = logLevel
	constants.LogConfig.LogFormat = style
}

func init() {
	constants.LogConfig.LogLevel = models.Debug
	constants.LogConfig.LogFormat = models.ColoredLogs
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

func Errorf(inp string, args ...any) {
	logs.Errorf(inp, args...)
}

func Infof(inp string, args ...any) {
	logs.Infof(inp, args...)
}

func Debugf(inp string, args ...any) {
	logs.Debugf(inp, args...)
}

func Warnf(inp string, args ...any) {
	logs.Warnf(inp, args...)
}

func ErrorWithTrace(args ...any) {
	logs.ErrorWithTrace(args...)
}
