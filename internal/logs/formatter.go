package logs

import (
	"encoding/json"
	"fmt"

	"github.com/KIVUOS1999/easyLogs/internal/constants"
	"github.com/KIVUOS1999/easyLogs/internal/models"
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func logf(level configs.LogLevel, args ...any) {
	if level > constants.LogConfig.LogLevel {
		return
	}

	inp, ok := args[0].(string)
	if !ok {
		return
	}

	params := args[1:]

	colorCode := lvlToColor(level)

	fmt.Printf(colorCode + generatePrefix(level))
	fmt.Printf(inp, params...)
	fmt.Println(constants.Reset)
}

func log(level configs.LogLevel, inp ...any) {
	if level == configs.ErrorWithTrace {
		printTrace(level, inp...)
		return
	}

	if level > constants.LogConfig.LogLevel {
		return
	}

	colorCode := lvlToColor(level)
	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	fmt.Println(colorCode + prefix + userLog + constants.Reset)
}

func printTrace(level configs.LogLevel, inp ...any) {
	colorCode := lvlToColor(level)
	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	fmt.Printf(colorCode+prefix+userLog+"\n%s"+constants.Reset+"\n", getStackTrace())
}

func jsonLog(level configs.LogLevel, inp ...any) {
	trace := ""
	if level == configs.ErrorWithTrace {
		trace = getStackTrace()
	}

	message := createLogString(inp...)
	loglevel := lvlToString(level)
	time := generateTimeString()

	log := models.JsonFormat{
		Message: message,
		Type:    loglevel,
		Trace:   trace,
		Time:    time,
	}

	jsonLog, _ := json.Marshal(log)
	fmt.Printf("%+v\n", string(jsonLog))
}

func jsonLogf(level configs.LogLevel, inp ...any) {
	input, ok := inp[0].(string)
	if !ok {
		return
	}

	args := inp[1:]
	userLog := fmt.Sprintf(input, args...)

	trace := ""
	if level == configs.ErrorWithTrace {
		trace = getStackTrace()
	}

	message := userLog
	loglevel := lvlToString(level)
	time := generateTimeString()

	log := models.JsonFormat{
		Message: message,
		Type:    loglevel,
		Trace:   trace,
		Time:    time,
	}

	jsonLog, _ := json.Marshal(log)
	fmt.Printf("%+v\n", string(jsonLog))
}

func formatDecider(level configs.LogLevel, formattedLogs bool, args ...any) {
	if constants.LogConfig.LogFormat == configs.JsonLogs {
		if formattedLogs {
			formatLog(level, jsonLogf, args...)
			return
		}

		formatLog(level, jsonLog, args...)
		return
	}

	if formattedLogs {
		formatLog(level, logf, args...)
		return
	}

	formatLog(level, log, args...)
}

func formatLog(level configs.LogLevel, handler LogHandler, args ...any) {
	handler(level, args...)
}

type LogHandler func(configs.LogLevel, ...any)
