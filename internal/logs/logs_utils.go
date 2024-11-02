package logs

import (
	"fmt"
	"strings"
	"time"

	constants "github.com/KIVUOS1999/easyLogs/internal"
)

func lvlToColor(level int) string {
	colorCode := constants.White
	switch level {
	case constants.Error:
		colorCode = constants.Red
	case constants.ErrorWithTrace:
		colorCode = constants.Red
	case constants.Info:
		colorCode = constants.Blue
	case constants.Debug:
		colorCode = constants.White
	case constants.Warn:
		colorCode = constants.Yellow
	}

	return colorCode
}

func generatePrefix(level int) string {
	levelName := "[ DEBUG ]\t:"

	switch level {
	case constants.Error:
		levelName = "[ ERROR ]\t:"
	case constants.Warn:
		levelName = "[ WARN ]\t:"
	case constants.Info:
		levelName = "[ INFO ]\t:"
	case constants.ErrorWithTrace:
		levelName = "[ TRACE ]\t:"
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

func logf(level int, inp string, args ...any) {
	colorCode := lvlToColor(level)

	fmt.Printf(colorCode + generatePrefix(level))
	fmt.Printf(inp, args...)
	fmt.Println(constants.Reset)
}

func log(level int, inp ...any) {
	colorCode := lvlToColor(level)

	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	if level == constants.ErrorWithTrace {
		fmt.Printf(colorCode+prefix+userLog+"%s"+constants.Reset+"\n", printStackTrace())
		return
	}

	fmt.Println(colorCode + prefix + userLog + constants.Reset)
}
