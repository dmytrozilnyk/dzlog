package dzlog

import (
	"context"
	"io"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/apsdehal/go-logger"
	"github.com/davecgh/go-spew/spew"
)

func New(opts ...OptsFunc) Dzlog {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	if os.Getenv(envLogLevel) != "" { //Get log level from environment variable LOG_LEVEL
		o.logLevel = stringToLogLevel(os.Getenv(envLogLevel))
	}

	return &customLog{opts: o}
}

func (l *customLog) GetTimeNow() time.Time {
	if l.opts.logLevel == DebugLevel {
		return time.Now()
	} else {
		return time.Time{}
	}
}

func (l *customLog) Debug(msg ...any) {
	if l.opts.logLevel == DebugLevel {
		apacheLog(l.opts.stdOut, l.opts.module, DebugLevel, generateMsg(msg), l.opts.logMaxLength, nil, l.opts.env)
	}
}

func (l *customLog) DebugTime(msg ...any) {
	if l.opts.logLevel == DebugLevel {
		apacheLog(l.opts.stdOut, l.opts.module, DebugLevel, generateMsg(msg), l.opts.logMaxLength, nil, l.opts.env)
	}
}

func (l *customLog) Info(msg ...any) {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel {
		apacheLog(l.opts.stdOut, l.opts.module, InfoLevel, generateMsg(msg), l.opts.logMaxLength, nil, l.opts.env)
	}
}

func (l *customLog) Warning(msg ...any) {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel || l.opts.logLevel == WarnLevel {
		apacheLog(l.opts.stdOut, l.opts.module, WarnLevel, generateMsg(msg), l.opts.logMaxLength, nil, l.opts.env)
	}
}

func (l *customLog) Error(msg ...any) {
	apacheLog(l.opts.stdOut, l.opts.module, ErrorLevel, generateMsg(msg), l.opts.logMaxLength, nil, l.opts.env)
}

// Context
func (l *customLog) DebugCtx(ctx context.Context, msg ...any) {
	if l.opts.logLevel == DebugLevel {
		apacheLog(l.opts.stdOut, l.opts.module, DebugLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
}

func (l *customLog) DebugTimeCtx(ctx context.Context, msg ...any) {
	if l.opts.logLevel == DebugLevel {
		apacheLog(l.opts.stdOut, l.opts.module, DebugLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
}

func (l *customLog) InfoCtx(ctx context.Context, msg ...any) {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel {
		apacheLog(l.opts.stdOut, l.opts.module, InfoLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
}

func (l *customLog) WarningCtx(ctx context.Context, msg ...any) {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel || l.opts.logLevel == WarnLevel {
		apacheLog(l.opts.stdOut, l.opts.module, WarnLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
}

func (l *customLog) ErrorCtx(ctx context.Context, msg ...any) {
	apacheLog(l.opts.stdOut, l.opts.module, ErrorLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
}

func (l *customLog) CalculateDifference(initial time.Time) time.Duration {
	if l.opts.logLevel == DebugLevel {
		return time.Since(initial)
	} else {
		return time.Duration(0)
	}
}

// return context
func (l *customLog) DebugReturnCtx(ctx context.Context, msg ...any) context.Context {
	if l.opts.logLevel == DebugLevel {
		apacheLog(l.opts.stdOut, l.opts.module, DebugLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
	return ctx
}

func (l *customLog) InfoReturnCtx(ctx context.Context, msg ...any) context.Context {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel {
		apacheLog(l.opts.stdOut, l.opts.module, InfoLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}
	return ctx
}

func (l *customLog) WarningReturnCtx(ctx context.Context, msg ...any) context.Context {
	if l.opts.logLevel == DebugLevel || l.opts.logLevel == InfoLevel || l.opts.logLevel == WarnLevel {
		apacheLog(l.opts.stdOut, l.opts.module, WarnLevel, generateMsg(msg), l.opts.logMaxLength, ctx, l.opts.env)
	}

	return ctx
}

func apacheLog(out io.Writer, module string, level LoggerLevel, msg string, logMaxLength int, ctx context.Context, env string) context.Context {
	logCustom, err := logger.New(module, out)
	if err != nil {
		panic(err)
	}
	logCustom.SetFormat("[%{time:2006-02-01 15:04:05.000}][%{module}%{message}")
	logCustom.SetLogLevel(logger.DebugLevel)

	pid := syscall.Getpid()
	//tid := syscall.Gettid()//Ubuntu
	tid := 0
	_, fileName, line, ok := runtime.Caller(3)
	slash := strings.LastIndex(fileName, "/")
	fileName = fileName[slash+1:]
	if !ok {
		fileName = "???"
		line = 0
	}
	var tracId = ""

	if ctx != nil {
		if ctx.Value(tracerId) != nil {
			tracId = ctx.Value(tracerId).(string)
		}
	}

	l := math.Ceil(float64(len(msg)) / float64(logMaxLength))
	length := int(l)
	for i := 0; i < length; i++ {
		var mat string
		if i == length-1 {
			lastBlock := i * logMaxLength
			mat = msg[lastBlock:]
		} else {
			first := i * logMaxLength
			last := first + logMaxLength
			mat = msg[first:last]
		}
		entries := strconv.Itoa(i+1) + "/" + strconv.Itoa(length)

		switch level {
		case ErrorLevel:
			logCustom.ErrorF(":%s][env:%s][pid:%d tid:%d][tracerId: %s][entries: %s][%s:%d] %s", level, env, pid, tid, tracId, entries, fileName, line, mat)
		case WarnLevel:
			logCustom.WarningF(":%s][env:%s][pid:%d tid:%d][tracerId: %s][entries: %s][%s:%d] %s", level, env, pid, tid, tracId, entries, fileName, line, mat)
		case InfoLevel:
			logCustom.InfoF(":%s][env:%s][pid:%d tid:%d][tracerId: %s][entries: %s][%s:%d] %s", level, env, pid, tid, tracId, entries, fileName, line, mat)
		case DebugLevel:
			logCustom.DebugF(":%s][env:%s][pid:%d tid:%d][tracer: %s][entries: %s][%s:%d] %s", level, env, pid, tid, tracId, entries, fileName, line, mat)
		default:
			logCustom.InfoF(":%s][env:%s][pid:%d tid:%d][tracerId: %s][entries: %s][%s:%d] %s", level, env, pid, tid, tracId, entries, fileName, line, mat)
		}
	}

	return ctx
}

func generateMsg(msg []any) string {

	msgString := ""

	allNil := true
	for _, arg := range msg {
		if arg != nil {
			allNil = false
			break
		}
	}

	if allNil {
		return msgString
	}

	for i, arg := range msg {
		argString := spew.Sprint(arg)
		msgString += argString
		if i < len(msg)-1 {
			msgString += " "
		}
	}

	return msgString
}
