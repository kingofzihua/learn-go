package logging

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Value

func init() {
	defaultLogger.Store(New(NewDefaultHandler(os.Stdout)))
}

// Default returns the default Logger.
func Default() *Logger { return defaultLogger.Load().(*Logger) }

// SetDefault makes l the default Logger.
// After this call, output from the log package's default Logger
// (as with [log.Print], etc.) will be logged at LevelInfo using l's Handler.
func SetDefault(l *Logger) {
	defaultLogger.Store(l)
}

type Logger struct {
	handler Handler
	// withAttrs
	attrs []interface{}
	// has Valuer
	hasValuer bool
	ctx       context.Context
}

func (l *Logger) With(args ...any) *Logger {
	l.attrs = append(l.attrs, args...)
	l.hasValuer = containsValuer(args)
	return l
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	l.ctx = ctx
	return l
}

// New creates a new Logger with the given non-nil Handler and a nil context.
func New(h Handler) *Logger {
	if h == nil {
		panic("nil Handler")
	}
	return &Logger{handler: h, ctx: context.Background()}
}

func (l *Logger) log(level Level, msg string, kvs ...interface{}) {
	keyVales := make([]interface{}, 0, len(l.attrs)+len(kvs))
	keyVales = append(keyVales, l.attrs...)
	if l.hasValuer {
		bindValues(l.ctx, keyVales)
	}
	keyVales = append(keyVales, kvs...)
	l.handler.Handle(level, msg, keyVales...)
}

func (l *Logger) Log(level Level, msg string, args ...any) {
	l.log(level, msg, args...)
}

// Debug logs at LevelDebug.
func (l *Logger) Debug(msg string, args ...any) {
	l.log(LevelDebug, msg, args...)
}

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.log(LevelInfo, msg, args...)
}

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.log(LevelWarn, msg, args...)
}

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any) {
	l.log(LevelError, msg, args...)
}

func Log(level Level, msg string, args ...any) {
	Default().log(level, msg, args...)
}

// Debug logs at LevelDebug.
func Debug(msg string, args ...any) {
	Default().Debug(msg, args...)
}
func Debugf(format string, a ...any) {
	Default().Debug(fmt.Sprintf(format, a...))
}

// Info logs at LevelInfo.
func Info(msg string, args ...any) {
	Default().Info(msg, args...)
}

// Infof logs at LevelInfo.
func Infof(format string, a ...any) {
	Default().Info(fmt.Sprintf(format, a...))
}

// Warn logs at LevelWarn.
func Warn(msg string, args ...any) {
	Default().Warn(msg, args...)
}

// Warnf logs at LevelWarn.
func Warnf(format string, a ...any) {
	Default().Warn(fmt.Sprintf(format, a...))
}

// Error logs at LevelError.
func Error(msg string, args ...any) {
	Default().Error(msg, args...)
}

// Errorf logs at LevelError.
func Errorf(format string, a ...any) {
	Default().Error(fmt.Sprintf(format, a...))
}

func With(l *Logger, kvs ...any) *Logger {
	attrs := make([]interface{}, 0, len(l.attrs)+len(kvs))
	attrs = append(attrs, l.attrs...)
	attrs = append(attrs, kvs...)
	return &Logger{handler: l.handler, attrs: attrs, hasValuer: containsValuer(attrs), ctx: l.ctx}
}

// WithContext returns a shallow copy of h with its context changed
// to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context, l *Logger) *Logger {
	return &Logger{handler: l.handler, attrs: l.attrs, hasValuer: l.hasValuer, ctx: ctx}
}
