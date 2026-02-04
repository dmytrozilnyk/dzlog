package dzlog

import (
	"io"
	"os"
)

const (
	typeLog          string = "APACHE"
	envLogLevel      string = "LOG_LEVEL"
	defaultMsgLength int    = 3000
	correlationId    string = "correlationId"
)

type customLog struct {
	opts opts
}

type opts struct {
	//Name of module that used Logger
	module string
	//Config Level. It can be specified using this variable, the global variable LOG_LEVEL or the FilePath configuration file.
	logLevel LoggerLevel
	//Maximum message length. Number of characters
	logMaxLength int
	//Log output
	stdOut io.Writer
	stdErr io.Writer
}

type OptsFunc func(*opts)

func defaultOpts() opts {
	return opts{
		module:       "Logger",
		logLevel:     InfoLevel,
		logMaxLength: defaultMsgLength,
		stdOut:       os.Stdout,
		stdErr:       os.Stderr,
	}
}

func LogLevel(logLevel LoggerLevel) OptsFunc {
	return func(opts *opts) {
		opts.logLevel = logLevel
	}
}

func LogModuleName(name string) OptsFunc {
	return func(opts *opts) {
		opts.module = name
	}
}

func LogMaxLength(numberOfChar int) OptsFunc {
	return func(opts *opts) {
		opts.logMaxLength = numberOfChar
		if numberOfChar <= 0 {
			opts.logMaxLength = defaultMsgLength
		}
	}
}
