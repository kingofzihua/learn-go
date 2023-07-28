package logging

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"sync"
)

// DefaultMessageKey default message key.
var DefaultMessageKey = "msg"

type Handler interface {
	Handle(level Level, msg string, args ...any)
}

type defaultHandler struct {
	log  *log.Logger
	pool *sync.Pool
}

// NewDefaultHandler new a std.log with writer.
func NewDefaultHandler(w io.Writer) Handler {
	return &defaultHandler{
		log: log.New(w, "", 0),
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (l *defaultHandler) Handle(level Level, msg string, kvs ...interface{}) {
	if (len(kvs) & 1) == 1 {
		kvs = append(kvs, "KEYVALS UNPAIRED")
	}
	buf := l.pool.Get().(*bytes.Buffer)
	buf.WriteString(level.String())
	_, _ = fmt.Fprintf(buf, " %s=%v", DefaultMessageKey, msg)
	for i := 0; i < len(kvs); i += 2 {
		_, _ = fmt.Fprintf(buf, " %s=%v", kvs[i], kvs[i+1])
	}
	_ = l.log.Output(4, buf.String()) //nolint:gomnd
	buf.Reset()
	l.pool.Put(buf)
}

func (l *defaultHandler) Close() error {
	return nil
}
