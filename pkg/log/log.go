package log

import (
	"github.com/KIVUOS1999/easyLogs/internal/logs"
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func Init(args ...any) {
	if len(args) >= 1 {
		if level, ok := args[0].(configs.LogLevel); ok {
			logs.LogConfig.LogLevel = level
		}
	}
	if len(args) >= 2 {

		if format, ok := args[1].(configs.LogFormat); ok {
			logs.LogConfig.LogFormat = format
		}
	}

	if len(args) >= 3 {
		if memStats, ok := args[2].(bool); ok {
			logs.LogConfig.EnableMemStats = memStats
		}
	}
}

func init() {
	logs.LogConfig.LogLevel = configs.Debug
	logs.LogConfig.LogFormat = configs.ColoredLogs
	logs.LogConfig.EnableMemStats = false
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
