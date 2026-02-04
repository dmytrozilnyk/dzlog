package dzlog

type LoggerLevel string

const (
	ErrorLevel LoggerLevel = "ERROR"
	WarnLevel  LoggerLevel = "WARN"
	InfoLevel  LoggerLevel = "INFO"
	DebugLevel LoggerLevel = "DEBUG"
)
