package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)
	fmt.Println("Lenght of array 1", len(colors))
	fmt.Println("Lenght of array 2", len(numbers))
}
