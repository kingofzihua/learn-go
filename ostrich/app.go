package ostrich

import (
	"context"
	"sync"
	"syscall"

	"github.com/Zemanta/gracefulshutdown"
	"github.com/Zemanta/gracefulshutdown/shutdownmanagers/posixsignal"
	"github.com/spf13/cobra"
)

type RunFunc func(basename string) error

type App struct {
	opts    options
	gs      *gracefulshutdown.GracefulShutdown
	ctx     context.Context
	cancel  func()
	mu      sync.Mutex
	runFunc RunFunc

	args cobra.PositionalArgs
	cmd  *cobra.Command
}

func New(opts ...Option) *App {
	o := options{
		gsm: posixsignal.NewPosixSignalManager(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT),
	}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())
	gs := gracefulshutdown.New()

	gs.AddShutdownManager(o.gsm)

	a := &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
		gs:     gs,
	}

	a.buildCommand()

	return a
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:   a.opts.name,
		Short: a.opts.name,
		Long:  a.opts.name,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}

	a.cmd = &cmd
}

func (a *App) Run() error {
	// start shutdown managers
	if err := a.gs.Start(); err != nil {
		return err
	}

	return a.cmd.Execute()
}
