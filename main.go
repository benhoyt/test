package main

import (
	"fmt"
)

// Barf some belarfen - tweaked more and more
func barf(n int) {
	fmt.Println("Goodbye, world - t3", n)
}

func main() {
	barf(42)
}
