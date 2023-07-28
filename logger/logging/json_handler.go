package logging

import (
	"context"
	"golang.org/x/exp/slog"
	"io"
	"runtime"
	"time"
)

var _ Handler = (*JSONHandler)(nil)

type JSONHandler struct {
	handler *slog.JSONHandler
}

func NewJSONHandler(w io.Writer) *JSONHandler {
	opts := &slog.HandlerOptions{
		AddSource: false,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// replace slog.TimeKey to empty
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			// replace slog.TimeKey to empty
			if a.Key == slog.SourceKey {
				return slog.Attr{}
			}

			// replace slog.MessageKey to local DefaultMessageKey
			if a.Key == slog.MessageKey {
				return slog.Attr{Key: DefaultMessageKey, Value: a.Value}
			}

			// replace slog.LevelKey to local LevelKey
			if a.Key == slog.LevelKey {
				return slog.Attr{Key: LevelKey, Value: a.Value}
			}
			return a
		},
	}

	return &JSONHandler{
		handler: slog.NewJSONHandler(w, opts),
	}
}

func (h *JSONHandler) Handle(level Level, msg string, args ...any) {
	var pc uintptr
	{
		var pcs [1]uintptr
		// skip [runtime.Callers, this function, this function's caller]
		runtime.Callers(3, pcs[:])
		pc = pcs[0]
	}
	r := slog.NewRecord(time.Now(), slog.Level(level), msg, pc)
	r.Add(args...)

	_ = h.handler.Handle(context.Background(), r)
}
