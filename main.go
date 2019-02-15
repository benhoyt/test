package main

import (
	"fmt"
)

// Barf some belarfen - tweaked
func barf(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	barf(42)
}
