package main

import (
	"fmt"
)

func main() {
	answer := 42
	var result string
	if answer < 0 {
		result = "Less than zero"
	} else if answer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)
}
