package logs

import (
	"encoding/json"
	"fmt"
	"path"

	"github.com/KIVUOS1999/easyLogs/internal/models"
	"github.com/KIVUOS1999/easyLogs/pkg/configs"
)

func logf(level configs.LogLevel, args ...any) {
	if level > LogConfig.LogLevel {
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
	fmt.Println(reset)
}

func log(level configs.LogLevel, inp ...any) {
	if level == configs.ErrorWithTrace {
		printTrace(level, inp...)
		return
	}

	if level > LogConfig.LogLevel {
		return
	}

	colorCode := lvlToColor(level)
	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	fmt.Println(colorCode + prefix + userLog + reset)
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

	formatStatsAndCallerDetails(&log)

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

	formatStatsAndCallerDetails(&log)

	jsonLog, _ := json.Marshal(log)
	fmt.Printf("%+v\n", string(jsonLog))
}

func formatStatsAndCallerDetails(input *models.JsonFormat) {
	callerDeatils := getCallerFuncName(skip)
	runTimeStats := getSystemStats()

	if callerDeatils != nil {
		input.CallerFileName = path.Base(callerDeatils.CallerFileName)
		input.CallerFunctionName = callerDeatils.CallerFunctionName
		input.LineNumber = callerDeatils.LineNumber
	}

	if runTimeStats != nil {
		input.HeapAlloc = runTimeStats.HeapAlloc
		input.StackAlloc = runTimeStats.StackAlloc
		input.TotalAlloc = runTimeStats.TotalAlloc
		input.SysAlloc = runTimeStats.SysAlloc
		input.NumGC = runTimeStats.NumGC
		input.NoGoRoutine = runTimeStats.NoGoRoutine
	}

}

func printTrace(level configs.LogLevel, inp ...any) {
	colorCode := lvlToColor(level)
	prefix := generatePrefix(level)
	userLog := createLogString(inp...)

	fmt.Printf(colorCode+prefix+userLog+"\n%s"+reset+"\n", getStackTrace())
}

func formatDecider(level configs.LogLevel, formattedLogs bool, args ...any) {
	if LogConfig.LogFormat == configs.JsonLogs {
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
