package for_loop

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		i += 1
	}
}
