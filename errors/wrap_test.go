package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestReadConfigWrap(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("error : %v", err)
	}
}

func TestReadConfigCause(t *testing.T) {
	_, err := ReadConfig()

	if err != nil {
		fmt.Printf("cause : %v\n", errors.Cause(err))
	}
}

func TestReadConfigPrintStack(t *testing.T) {
	_, err := ReadConfig()

	if err != nil {
		fmt.Printf("error : %+v\n", err)
	}
}

func TestErrorf(t *testing.T) {
	var ErrPermission = errors.New("permission err")

	err := fmt.Errorf("access denied: %w", ErrPermission)

	// ...
	if errors.Is(err, ErrPermission) { // true
		fmt.Println("is permission error")
	}
	// fmt.Printf("%+v", errors.WithStack(err))
}
