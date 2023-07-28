package log

import (
	"context"
	"fmt"
	"log"
)

// DefaultMessageKey default message key.
var DefaultMessageKey = "msg"

// DefaultLogger is default logger.
var DefaultLogger = NewStdLogger(log.Writer())

var global = DefaultLogger

// GetLogger returns global logger appliance as logger in current process.
func GetLogger() Logger {
	return global
}

// SetLogger should be called before any other log call.
// And it is NOT THREAD SAFE.
func SetLogger(l Logger) {
	global = l
}

// Log Print log by level and keyvals.
func Log(level Level, msg string, kvs ...interface{}) {
	global.Log(level, msg, kvs...)
}

// Context with context logger.
func Context(ctx context.Context) Logger {
	return WithContext(ctx, global)
}

// Debug logs a message at debug level.
func Debug(msg string, kvs ...interface{}) {
	global.Log(LevelDebug, msg, kvs...)
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	global.Log(LevelDebug, fmt.Sprintf(format, a...))
}

// Info logs a message at info level.
func Info(msg string, kvs ...interface{}) {
	global.Log(LevelInfo, msg, kvs...)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	global.Log(LevelInfo, fmt.Sprintf(format, a...))
}

// Warn logs a message at warn level.
func Warn(msg string, kvs ...interface{}) {
	global.Log(LevelWarn, msg, kvs...)
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	global.Log(LevelWarn, fmt.Sprintf(format, a...))
}

// Error logs a message at error level.
func Error(msg string, kvs ...interface{}) {
	global.Log(LevelError, msg, kvs...)
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	global.Log(LevelError, fmt.Sprintf(format, a...))
}
