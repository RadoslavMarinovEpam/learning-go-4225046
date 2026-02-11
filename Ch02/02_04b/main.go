package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	sum := i1 + i2 + i3
	fmt.Println("ISUM: ", sum)

	f1, f2, f3 := 12.5, 45.6, 68.9
	sumf := f1 + f2 + f3
	fmt.Println("FSUM: ", sumf)
	total := float64(i1) + f1
	fmt.Println("Total: ", total)
}
