package logging

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Valuer is returns a log value.
type Valuer func(ctx context.Context) interface{}

// Value return the function value.
func Value(ctx context.Context, val interface{}) interface{} {
	if v, ok := val.(Valuer); ok {
		return v(ctx)
	}
	return val
}

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) Valuer {
	return func(ctx context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return file[:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}

func bindValues(ctx context.Context, kvs []interface{}) {
	for i := 1; i < len(kvs); i += 2 {
		if v, ok := kvs[i].(Valuer); ok {
			kvs[i] = v(ctx)
		}
	}
}

// containsValuer Check kvs contains Valuer
func containsValuer(kvs []interface{}) bool {
	for i := 1; i < len(kvs); i += 2 {
		if _, ok := kvs[i].(Valuer); ok {
			return true
		}
	}
	return false
}
