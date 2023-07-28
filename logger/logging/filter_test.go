package logging

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTextFilterBufferLogger(opts ...FilterOption) (*Logger, *bytes.Buffer) {
	w := new(bytes.Buffer)
	handler := NewDefaultHandler(w)
	filter := NewFilter(handler, opts...)
	return New(filter), w
}

func TestNewFilter(t *testing.T) {
	log, _ := newTextFilterBufferLogger()
	assert.NotNil(t, log)
}

func TestFilterHandler_FilterLevel(t *testing.T) {
	log, w := newTextFilterBufferLogger(FilterLevel(LevelError))

	log.Log(LevelInfo, "info")

	assert.Empty(t, w.String())
}

func TestFilterHandler_FilterKey(t *testing.T) {
	log, w := newTextFilterBufferLogger(FilterKey("name", "value"))

	log.Log(LevelInfo, "info", "name", "value", "value", "v")

	assert.Equal(t, w.String(), fmt.Sprintf("INFO msg=info name=%s value=%s\n", fuzzyStr, fuzzyStr))
}

func TestFilterHandler_FilterValue(t *testing.T) {
	log, w := newTextFilterBufferLogger(FilterValue("name", "value"))

	log.Log(LevelInfo, "info", "name", "value", "value", "v")

	assert.Equal(t, w.String(), fmt.Sprintf("INFO msg=info name=%s value=%s\n", fuzzyStr, "v"))
}
