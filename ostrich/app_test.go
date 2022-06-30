package ostrich

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	app := New(
		ID("0705"),
		Name("ostrich"),
	)
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
