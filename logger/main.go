package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/exp/slog"
	"io"
	"net"
	"os"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := file.WriteString("content\n"); err != nil {
		panic(err)
	}

	slogs(file)
	kraots(file)
	if _, err := file.WriteString("content\n"); err != nil {
		panic(err)
	}
}

func kraots(w io.Writer) {

	logger := log.NewStdLogger(w)

	logger = log.With(logger,
		"time", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"user", log.Valuer(func(ctx context.Context) any {
			u, ok := ctx.Value("name").(string)
			if ok {
				return u
			}
			return ""
		}),
	)

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")
	logger = log.WithContext(ctx, logger)

	log.SetLogger(logger)

	log.Infof("hello %s", "world")
	log.Context(ctx).Infow("msg", "warn msg", "domain", "baidu")
	log.Error("oops", "err", net.ErrClosed, "status", 500)
}

func slogs(w io.Writer) {
	l := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time from the output for predictable test output.
			if a.Key == "name" {
				return slog.Attr{}
			}
			return a
		},
	}))

	slog.SetDefault(l.With("name", "default"))

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")

	slog.InfoContext(ctx, "warn msg", "domain", "baidu")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
}
