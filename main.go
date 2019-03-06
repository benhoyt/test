package main

import (
	"fmt"
)

// Output some foo
func foo(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	foo(42)
}
