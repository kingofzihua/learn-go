package logging

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJSONHandler(t *testing.T) {
	log, _ := newJsonBufferLogger()

	assert.NotNil(t, log)
}

func newJsonBufferLogger() (*Logger, *bytes.Buffer) {
	w := new(bytes.Buffer)
	return New(NewJSONHandler(w)), w
}

func TestJSONHandler_Handle(t *testing.T) {
	log, w := newJsonBufferLogger()

	log.Log(LevelInfo, "handler")

	assert.Equal(t, w.String(), fmt.Sprintln(`{"level":"INFO","msg":"handler"}`))
}

func TestJSONHandler_HandleWith(t *testing.T) {
	log, w := newJsonBufferLogger()

	log.With("name", "kingofzihua").Log(LevelInfo, "handler")

	assert.Equal(t, w.String(), fmt.Sprintln(`{"level":"INFO","msg":"handler","name":"kingofzihua"}`))
}

func TestJSONHandler_HandleAttr(t *testing.T) {
	log, w := newJsonBufferLogger()

	log.Log(LevelInfo, "handler", "name", "kingofzihua")

	assert.Equal(t, w.String(), fmt.Sprintln(`{"level":"INFO","msg":"handler","name":"kingofzihua"}`))
}

func TestJSONHandler_HandleValuer(t *testing.T) {
	log, w := newJsonBufferLogger()

	log = log.With("name", Valuer(func(ctx context.Context) interface{} {
		return ctx.Value("name").(string)
	}))

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")

	log.WithContext(ctx).Log(LevelInfo, "handler")

	assert.Equal(t, w.String(), fmt.Sprintln(`{"level":"INFO","msg":"handler","name":"kingofzihua"}`))
}
