package main

import (
	"fmt"
)

func barf(n int) {
	fmt.Println("GOODBYE, world", n)
}

func main() {
	barf(42)
}
