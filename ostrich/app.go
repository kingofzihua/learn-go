package ostrich

import (
	"context"
	"sync"
	"os"

	"github.com/Zemanta/gracefulshutdown"
)

type App struct {
	opts   options
	gs     *gracefulshutdown.GracefulShutdown
	ctx    contex.Context
	cancel func()
	mu     sync.Mutex
}

func (a *App) Run() error {

}

func New(opts ...Option) *App {
	o := options{
		gsm: posixsignal.NewPosixSignalManager([]os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}),
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())

	gs := gracefulshutdown.New()

	gs.AddShutdownManager(gsm)

	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
		gs:     gs,
	}
}
