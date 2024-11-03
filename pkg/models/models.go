package models

type Logger struct{}

type LoggerConfig struct {
	LogLevel  LogLevel
	LogFormat LogFormat
}

type LogLevel int

const (
	Error LogLevel = iota
	Warn
	Info
	Debug
	ErrorWithTrace
)

type LogFormat int

const (
	ColoredLogs LogFormat = iota
	JsonLogs
)
