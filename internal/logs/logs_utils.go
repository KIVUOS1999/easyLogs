package logs

import (
	"fmt"
	"strings"
	"time"

	"github.com/KIVUOS1999/easyLogs/internal/constants"
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func lvlToString(level configs.LogLevel) string {
	var levelNameMap = map[configs.LogLevel]string{
		configs.Info:           "Info",
		configs.Debug:          "Debug",
		configs.Error:          "Error",
		configs.Warn:           "Warn",
		configs.ErrorWithTrace: "Trace",
	}

	return levelNameMap[level]
}

func lvlToColor(level configs.LogLevel) string {
	colorCode := constants.White
	switch level {
	case configs.Error:
		colorCode = constants.Red
	case configs.ErrorWithTrace:
		colorCode = constants.Red
	case configs.Info:
		colorCode = constants.Blue
	case configs.Debug:
		colorCode = constants.White
	case configs.Warn:
		colorCode = constants.Yellow
	}

	return colorCode
}

func generatePrefix(level configs.LogLevel) string {
	levelName := "[ DEBUG ]\t"

	switch level {
	case configs.Error:
		levelName = "[ ERROR ]\t"
	case configs.Warn:
		levelName = "[ WARN ]\t"
	case configs.Info:
		levelName = "[ INFO ]\t"
	case configs.ErrorWithTrace:
		levelName = "[ TRACE ]\t"
	}

	currentTime := generateTimeString()
	return currentTime + "\t" + levelName
}

func generateTimeString() string {
	utcTime := time.Now().Local()
	structuredTime := utcTime.Format("2006-01-02 15:04:05")

	return structuredTime
}

func createLogString(parts ...any) string {
	var stringParts []string
	for _, part := range parts {
		stringParts = append(stringParts, fmt.Sprintf("%v", part))
	}

	logString := strings.Join(stringParts, " ")
	return logString
}
