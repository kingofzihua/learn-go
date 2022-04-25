package errors

import (
	"testing"
)

func TestNewCustomError(t *testing.T) {
	ce := NewCustomError("EOF")

	if ce == NewCustomError("EOF") {
		t.Error("Custom Error Equal")
	}
}
