package log

import (
	"context"
)

// A Logger represents a logger.
type Logger interface {
	// Debug logs a message at debug level.
	Debug(...any)
	// Debugf logs a message at debug level.
	Debugf(string, ...any)
	// Debugw logs a message at debug level.
	Debugw(string, ...any)
	// Info logs a message at info level.
	Info(...any)
	// Infof logs a message at info level.
	Infof(string, ...any)
	// Infow logs a message at info level.
	Infow(string, ...any)
	// Error logs a message at error level.
	Error(...any)
	// Errorf logs a message at error level.
	Errorf(string, ...any)
	// Errorw logs a message at error level.
	Errorw(string, ...any)
	// Fatal logs a message at error level.
	Fatal(...any)
	// Fatalf logs a message at error level.
	Fatalf(string, ...any)
	// Fatalw logs a message at error level.
	Fatalw(string, ...any)
	// WithContext returns a new logger with the given context.
	WithContext(ctx context.Context) Logger
}
