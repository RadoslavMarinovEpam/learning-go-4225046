package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println(AddValues(2, 5))
	fmt.Println(addAllValues(1, 2, 3, 4, 5))
}

func doSomething() {
	fmt.Println("Doing something")
}

func AddValues(a, b int) int {
	return a + b
}

func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	count := len(values)
	average := float64(sum) / float64(count)
	return sum, count, average
}
