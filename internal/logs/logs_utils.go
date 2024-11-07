package logs

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/KIVUOS1999/easyLogs/internal/models"
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
	colorCode := white
	switch level {
	case configs.Error:
		colorCode = red
	case configs.ErrorWithTrace:
		colorCode = red
	case configs.Info:
		colorCode = blue
	case configs.Debug:
		colorCode = white
	case configs.Warn:
		colorCode = yellow
	}

	return colorCode
}

func generatePrefix(level configs.LogLevel) string {
	levelName := "[ DEBUG ]"

	switch level {
	case configs.Error:
		levelName = "[ ERROR ]"
	case configs.Warn:
		levelName = "[ WARN  ]"
	case configs.Info:
		levelName = "[ INFO  ]"
	case configs.ErrorWithTrace:
		levelName = "[ TRACE ]"
	}

	currentTime := generateTimeString()
	funcName := getCallerFuncName(skip)
	return currentTime + " " + levelName + " " + funcName.CallerFunctionName + "()#" + strconv.Itoa(funcName.LineNumber) + "\t"
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

func getCallerFuncName(skip int) *models.JsonFormat {
	pc, file, no, ok := runtime.Caller(skip)
	if ok {
		funcObj := runtime.FuncForPC(pc)
		funcName := funcObj.Name()

		return &models.JsonFormat{
			CallerFileName:     file,
			CallerFunctionName: funcName,
			LineNumber:         no,
		}
	}

	return nil
}

func getSystemStats() *models.JsonFormat {
	if !LogConfig.EnableMemStats {
		return nil
	}

	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)

	heapAlloc := math.Round(float64(memStats.Alloc)/1024/1024*100) / 100
	stackAlloc := math.Round(float64(memStats.StackInuse)/1024/1024*100) / 100
	totalAlloc := math.Round(float64(memStats.TotalAlloc)/1024/1024*100) / 100
	sysAlloc := math.Round(float64(memStats.Sys)/1024/1024*100) / 100
	numGC := int(memStats.NumGC)
	numGoRoutine := runtime.NumGoroutine()

	return &models.JsonFormat{
		HeapAlloc:   &heapAlloc,
		StackAlloc:  &stackAlloc,
		TotalAlloc:  &totalAlloc,
		SysAlloc:    &sysAlloc,
		NumGC:       &numGC,
		NoGoRoutine: &numGoRoutine,
	}
}
