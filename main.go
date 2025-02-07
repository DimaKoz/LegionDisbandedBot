package main

import (
	"fmt"
)

func main() {
	s := "gopher"
	// permit
	fmt.Printf("Hello and welcome, %s!\n", s) //nolint:forbidigo

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i) //nolint:forbidigo
	}
}
