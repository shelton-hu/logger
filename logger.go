package logger

import (
	"context"
	"io"
	"os"
)

type Logger struct {
	level         Level
	output        io.Writer
	hideCallstack bool
	depth         int
	suffixFn      func(ctx context.Context) string
}

var logger = newLogger()

func newLogger() *Logger {
	return &Logger{
		level:  InfoLevel,
		output: os.Stdout,
		depth:  4,
		suffixFn: func(ctx context.Context) string {
			return ""
		},
	}
}

func SetLevel(level Level) {
	logger.level = level
}

func SetLevelByString(level string) {
	logger.level = StringToLevel(level)
}

func SetOutput(output io.Writer) {
	logger.output = output
}

func SetHideCallSatck(hideCallStack bool) {
	logger.hideCallstack = hideCallStack
}

func SetDepth(depth int) {
	logger.depth = depth
}

func SetSuffixFn(suffixFn func(ctx context.Context) string) {
	logger.suffixFn = suffixFn
}

func Debug(ctx context.Context, format string, v ...interface{}) {
	logger.logf(ctx, DebugLevel, format, v...)
}

func Info(ctx context.Context, format string, v ...interface{}) {
	logger.logf(ctx, InfoLevel, format, v...)
}

func Warn(ctx context.Context, format string, v ...interface{}) {
	logger.logf(ctx, WarnLevel, format, v...)
}

func Error(ctx context.Context, format string, v ...interface{}) {
	logger.logf(ctx, ErrorLevel, format, v...)
}

func Critical(ctx context.Context, format string, v ...interface{}) {
	logger.logf(ctx, CriticalLevel, format, v...)
}
