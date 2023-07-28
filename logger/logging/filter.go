package logging

const fuzzyStr = "***"

// FilterOption is filter option.
type FilterOption func(*FilterHandler)

// FilterLevel with filter level.
func FilterLevel(level Level) FilterOption {
	return func(opts *FilterHandler) {
		opts.level = level
	}
}

// FilterKey with filter key.
func FilterKey(key ...string) FilterOption {
	return func(o *FilterHandler) {
		for _, v := range key {
			o.key[v] = struct{}{}
		}
	}
}

// FilterValue with filter value.
func FilterValue(value ...string) FilterOption {
	return func(o *FilterHandler) {
		for _, v := range value {
			o.value[v] = struct{}{}
		}
	}
}

type FilterHandler struct {
	handler Handler
	level   Level
	key     map[interface{}]struct{}
	value   map[interface{}]struct{}
	filter  Filter
}

// NewFilter new a logger filter.
func NewFilter(handler Handler, opts ...FilterOption) Handler {
	f := FilterHandler{
		handler: handler,
		level:   LevelInfo,
		key:     make(map[interface{}]struct{}),
		value:   make(map[interface{}]struct{}),
	}
	for _, o := range opts {
		o(&f)
	}
	return &f
}

type Filter func(level Level, kvs ...interface{}) bool

func (f *FilterHandler) Handle(level Level, msg string, kvs ...interface{}) {
	// level filter
	if level < f.level {
		return
	}

	// custom filter
	if f.filter != nil && (f.filter(level, kvs...)) {
		return
	}

	// key value filter
	if len(f.key) > 0 || len(f.value) > 0 {
		for i := 0; i < len(kvs); i += 2 {
			v := i + 1
			if v >= len(kvs) {
				continue
			}
			if _, ok := f.key[kvs[i]]; ok {
				kvs[v] = fuzzyStr
			}
			if _, ok := f.value[kvs[v]]; ok {
				kvs[v] = fuzzyStr
			}
		}
	}

	f.handler.Handle(level, msg, kvs...)
}
