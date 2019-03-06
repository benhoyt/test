package main

import (
	"fmt"
)

// Barf some belarfen - tweaked more and more
func foo(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	foo(42)
}
