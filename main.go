package main

import (
	"fmt"
)

func foo(n int) {
	fmt.Println("GOODBYE, world", n)
}

func main() {
	foo(42)
}
