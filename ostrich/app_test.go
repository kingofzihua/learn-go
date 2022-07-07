package ostrich

import (
	"fmt"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := New("ostrich", "", WithRunFunc(func(basename string) error {
		fmt.Println("app ostrich run is running!")
		return nil
	}))
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
