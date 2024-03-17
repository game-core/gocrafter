package logger

type LogLevel int32

const (
	Error   LogLevel = 0
	Warning LogLevel = 1
	Info    LogLevel = 2
	Debug   LogLevel = 3
	Success LogLevel = 4
)
