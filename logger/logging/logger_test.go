package logging

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func newBufferLogger() (*Logger, *bytes.Buffer) {
	w := new(bytes.Buffer)
	return New(NewDefaultHandler(w)), w
}

func setBufferStdLoggerToDefault() (*Logger, *bytes.Buffer) {
	logger, w := newBufferLogger()
	SetDefault(logger)
	return logger, w
}

func TestDefault(t *testing.T) {
	logger, _ := newBufferLogger()
	SetDefault(logger)
	assert.Equal(t, logger, Default())
}

func TestNew(t *testing.T) {
	log := New(NewDefaultHandler(os.Stdout))

	assert.NotNil(t, log)

	assert.Panics(t, func() {
		New(nil)

	})
}

func TestLogger_With(t *testing.T) {
	log, w := newBufferLogger()

	log.With("name", "kingofzihua").With("age", 18).Info("with")

	assert.Equal(t, w.String(), fmt.Sprintln("INFO msg=with name=kingofzihua age=18"))
}

func TestLogger_WithValuer(t *testing.T) {
	log, w := newBufferLogger()

	var name Valuer = func(ctx context.Context) interface{} {
		return ctx.Value("name")
	}

	log.With("name", name)

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")

	log.WithContext(ctx).Info("context")

	assert.Equal(t, w.String(), fmt.Sprintln("INFO msg=context name=kingofzihua"))
}

func TestLogger_Log(t *testing.T) {
	log, w := newBufferLogger()

	log.Log(LevelDebug, "debug")

	assert.Equal(t, w.String(), fmt.Sprintln("DEBUG msg=debug"))
}

func TestLogger_LogAttr(t *testing.T) {
	log, w := newBufferLogger()

	log.Log(LevelDebug, "debug", "name", "kingofzihua")

	assert.Equal(t, w.String(), fmt.Sprintln("DEBUG msg=debug name=kingofzihua"))
}

func TestGlobal(t *testing.T) {
	logger, w := setBufferStdLoggerToDefault()

	Log(LevelInfo, "info")
	Debug("debug")
	Debugf("user:%s", "kingofzihua")
	Info("info")
	Infof("user:%s", "kingofzihua")

	Warn("warn")
	Warnf("user:%s", "kingofzihua")
	Error("error")
	Errorf("user:%s", "kingofzihua")

	str := `INFO msg=info
DEBUG msg=debug
DEBUG msg=user:kingofzihua
INFO msg=info
INFO msg=user:kingofzihua
WARN msg=warn
WARN msg=user:kingofzihua
ERROR msg=error
ERROR msg=user:kingofzihua
INFO msg=with author=kingofzihua
INFO msg=ctx requestId=112358
INFO msg=ctx requestId=112358
`
	ctx := context.WithValue(context.Background(), "request-id", "112358")
	req := Valuer(func(ctx context.Context) interface{} {
		return ctx.Value("request-id")
	})
	With(logger, "author", "kingofzihua").Info("with")
	WithContext(ctx, logger).With("requestId", req).Info("ctx")
	// double
	WithContext(ctx, logger).With("requestId", req).Info("ctx")

	assert.Equal(t, w.String(), str)
}
