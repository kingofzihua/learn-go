package main

import (
	"context"
	"github.com/kingofzihua/learn-go/logger/kratos"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	//slogs()
	kraots()
}

func kraots() {
	logger := kratos.NewStdLogger(os.Stdout)

	logger = kratos.With(logger,
		"time", kratos.DefaultTimestamp,
		"caller", kratos.DefaultCaller,
		"user", kratos.Valuer(func(ctx context.Context) any {
			u, ok := ctx.Value("name").(string)
			if ok {
				return u
			}
			return ""
		}),
	)

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")
	logger = kratos.WithContext(ctx, logger)

	SetDefault(logger)

	Infof("hello %s", "world")
}

func slogs() {
	l := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	slog.SetDefault(l.With("name", func(ctx context.Context) string {
		return ctx.Value("name").(string)
	}))

	ctx := context.WithValue(context.Background(), "name", "kingofzihua")

	slog.InfoContext(ctx, "warn msg")
}
