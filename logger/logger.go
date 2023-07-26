package main

import (
	"fmt"
	"github.com/kingofzihua/learn-go/logger/kratos"
	"os"
	"sync/atomic"
)

var defaultLogger atomic.Value

func init() {
	logger := kratos.NewStdLogger(os.Stdout)
	logger = kratos.With(logger,
		"ts", kratos.DefaultTimestamp,
		"caller", kratos.DefaultCaller,
	)
	defaultLogger.Store(logger)
}

// Default returns the default Logger.
func Default() kratos.Logger { return defaultLogger.Load().(kratos.Logger) }

// SetDefault makes l the default Logger.
// After this call, output from the log package's default Logger
// (as with [log.Print], etc.) will be logged at LevelInfo using l's Handler.
func SetDefault(l kratos.Logger) {
	defaultLogger.Store(l)
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	_ = Default().Log(kratos.LevelInfo, "msg", fmt.Sprintf(format, a...))
	//Default().Log(LevelInfo, h.msgKey, fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func Infow(keyvals ...interface{}) {
	_ = Default().Log(kratos.LevelInfo, keyvals...)
}
