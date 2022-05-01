package errors

import (
	"github.com/pkg/errors"
	"testing"
)

func TestNewCustomError(t *testing.T) {
	ce := NewCustomError("EOF")
	ce2 := errors.New("EOF")

	if ce == ce2 {
		t.Error("Custom Error Equal")
	}
	if errors.Is(ce, ce2) {
		t.Error("Custom Error Equal")
	}
}
