package main

import (
	"fmt"
)

func main() {

	aint := 42
	var p *int = &aint

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Printf("p: %v\n", *p)
	}

	value1 := 42.13
	pointer1 := &value1
	*pointer1 = 43.14
	if pointer1 == nil {
		fmt.Println("pointer1 is nil")
	} else {
		fmt.Printf("pointer1: %v\n", *pointer1)
	}
}
