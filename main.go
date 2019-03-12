package main

import (
	"fmt"
)

// Output some foo pretty please
func foo(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	hello("world")
	thing()
	foo(24)
}
