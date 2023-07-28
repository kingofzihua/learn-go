package log

import "context"

// Logger is a logger interface.
type Logger interface {
	Log(level Level, msg string, args ...any)
}

type LoggerHandler interface {
	Handle(level Level, msg string, args ...any)
}

type logger struct {
	logger Logger
	// withAttrs
	attrs []interface{}
	// has Valuer
	hasValuer bool
	ctx       context.Context
}

func (l *logger) Log(level Level, msg string, kvs ...interface{}) {
	keyVales := make([]interface{}, 0, len(l.attrs)+len(kvs))
	keyVales = append(keyVales, l.attrs...)
	if l.hasValuer {
		bindValues(l.ctx, keyVales)
	}
	keyVales = append(keyVales, kvs...)
	l.logger.Log(level, msg, keyVales...)
}

// With is with logger fields.
func With(l Logger, kvs ...interface{}) Logger {
	c, ok := l.(*logger)
	if !ok {
		return &logger{logger: l, attrs: kvs, hasValuer: containsValuer(kvs), ctx: context.Background()}
	}

	attrs := make([]interface{}, 0, len(c.attrs)+len(kvs))
	attrs = append(attrs, c.attrs...)
	attrs = append(attrs, kvs...)

	return &logger{logger: c.logger, attrs: kvs, hasValuer: containsValuer(kvs), ctx: c.ctx}
}

// WithContext returns a shallow copy of l with its context changed to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context, l Logger) Logger {
	c, ok := l.(*logger)
	if ok {
		return &logger{logger: c.logger, attrs: c.attrs, hasValuer: c.hasValuer, ctx: c.ctx}
	}
	return &logger{logger: l, ctx: c.ctx}
}
