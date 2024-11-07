package models

import "github.com/KIVUOS1999/easyLogs/pkg/configs"

type Logger struct{}

type LoggerConfig struct {
	LogLevel       configs.LogLevel
	LogFormat      configs.LogFormat
	EnableMemStats bool
}

type JsonFormat struct {
	Type    string `json:"type"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Trace   string `json:"trace,omitempty"`

	CallerFileName     string `json:"caller_file"`
	CallerFunctionName string `json:"caller_func_name"`
	LineNumber         int    `json:"line_number"`

	HeapAlloc   *float64 `json:"current_heap_alloc_MB,omitempty"`
	StackAlloc  *float64 `json:"current_stack_alloc_MB,omitempty"`
	TotalAlloc  *float64 `json:"total_alloc_MB,omitempty"`
	SysAlloc    *float64 `json:"sys_alloc_MB,omitempty"`
	NumGC       *int     `json:"total_garbage_collected,omitempty"`
	NoGoRoutine *int     `json:"current_go_routine,omitempty"`
}
