package models

import "github.com/KIVUOS1999/easyLogs/pkg/configs"

type Logger struct{}

type LoggerConfig struct {
	LogLevel  configs.LogLevel
	LogFormat configs.LogFormat
}

type JsonFormat struct {
	Type    string `json:"type"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Trace   string `json:"trace,omitempty"`
}
