package main

import (
	"fmt"
)

func bar(n int) {
	fmt.Println("GOODBYE, world", n)
}

func main() {
	bar(42)
}
