package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an slice
	// var colors = []string{"Red", "Green", "Blue"}
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "cyan", "yellow", "lime")
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)

}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
