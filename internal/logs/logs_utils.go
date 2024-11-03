package logs

import (
	"fmt"
	"strings"
	"time"

	constants "github.com/KIVUOS1999/easyLogs/internal"
	"github.com/KIVUOS1999/easyLogs/pkg/models"
)

func lvlToColor(level models.LogLevel) string {
	colorCode := constants.White
	switch level {
	case models.Error:
		colorCode = constants.Red
	case models.ErrorWithTrace:
		colorCode = constants.Red
	case models.Info:
		colorCode = constants.Blue
	case models.Debug:
		colorCode = constants.White
	case models.Warn:
		colorCode = constants.Yellow
	}

	return colorCode
}

func generatePrefix(level models.LogLevel) string {
	levelName := "[ DEBUG ]\t"

	switch level {
	case models.Error:
		levelName = "[ ERROR ]\t"
	case models.Warn:
		levelName = "[ WARN ]\t"
	case models.Info:
		levelName = "[ INFO ]\t"
	case models.ErrorWithTrace:
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

func logf(level models.LogLevel, inp string, args ...any) {
	if level > constants.LogConfig.LogLevel {
		return
	}

	colorCode := lvlToColor(level)

	fmt.Printf(colorCode + generatePrefix(level))
	fmt.Printf(inp, args...)
	fmt.Println(constants.Reset)
}

func log(level models.LogLevel, inp ...any) {
	colorCode := lvlToColor(level)
	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	if level == models.ErrorWithTrace {
		fmt.Printf(colorCode+prefix+userLog+"\n%s"+constants.Reset+"\n", printStackTrace())
		return
	}

	if level > constants.LogConfig.LogLevel {
		return
	}

	fmt.Println(colorCode + prefix + userLog + constants.Reset)
}
