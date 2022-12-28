package main

import "fmt"

//go:noinline
func add(a, b int) int {
	return a + b
}

func main() {
	sum := add(1, 2)

	fmt.Println(sum)
}
