package main

import (
	"fmt"
)

// Barf some belarfen
func barf(n int) {
	fmt.Println("GOODBYE, world", n)
}

func main() {
	barf(42)
}
