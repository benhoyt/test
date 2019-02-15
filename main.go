package main

import (
	"fmt"
)

// Barf some belarfen - tweaked
func barf(n int) {
	fmt.Println("GOODBYE, world", n)
}

func main() {
	barf(42)
}
