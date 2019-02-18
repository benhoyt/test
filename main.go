package main

import (
	"fmt"
)

// Barf some belarfen - tweaked more and more
func barf(n int) {
	fmt.Println("Goodbye, world", n, 7)
}

func main() {
	barf(42)
}
