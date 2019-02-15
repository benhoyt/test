package main

import (
	"fmt"
)

func bar(n int) {
	fmt.Println("Goodbye, world", n)
}

func main() {
	bar(42)
}
