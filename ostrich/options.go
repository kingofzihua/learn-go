package ostrich

import (
	"github.com/Zemanta/gracefulshutdown"
)

type Option func(o *options)

type options struct {
	id       string
	name     string
	version  string
	metadata map[string]string
	gsm      gracefulshutdown.ShutdownManager
}

func ID(id string) Option {
	return func(o *options) { o.id = id }
}

func Name(name string) Option {
	return func(o *options) { o.name = name }
}

func Version(version string) Option {
	return func(o *options) { o.version = version }
}

func Metadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

func ShutdownManager(gsm gracefulshutdown.ShutdownManager) Option {
	return func(o *options) { o.gsm = gsm }
}
