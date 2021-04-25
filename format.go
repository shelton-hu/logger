package logger

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"
)

func (logger *Logger) formatOutput(ctx context.Context, level Level, output string) string {
	now := time.Now().Format("2006-01-02 15:04:05.99999")

	l := fmt.Sprintf(`-%s-`, level.String())
	output = strings.NewReplacer("\r", "\\r", "\n", "\\n").Replace(output)

	suffix := logger.suffixFn(ctx)

	if logger.hideCallstack {
		return fmt.Sprintf("%-25s %-10s %s %s",
			now, strings.ToUpper(l), output, suffix)
	} else {
		_, file, line, ok := runtime.Caller(logger.depth)
		if !ok {
			file = "???"
			line = 0
		}
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
		return fmt.Sprintf("%-25s %-10s %s (%s:%d)%s",
			now, strings.ToUpper(l), output, file, line, suffix)
	}
}

func (logger *Logger) logf(ctx context.Context, level Level, format string, args ...interface{}) {
	if logger.level < level {
		return
	}
	fmt.Fprintln(logger.output, logger.formatOutput(ctx, level, fmt.Sprintf(format, args...)))
}
