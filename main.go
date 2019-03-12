package main

import (
	"fmt"
)

// Output some foo pretty please
func foo(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	hi("world!")
	foo(24)
}
