package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	anumber := 42
	strLen, err := fmt.Println(str1, str2, str3)
	fmt.Println("The value is: ", anumber)

	if err == nil {
		fmt.Println("The length of the string is: ", strLen)
	} else {
		fmt.Println("An error occurred: ", err)
	}
	fmt.Printf("Value of number: %v\n", strLen)
	fmt.Printf("Data type: %T\n", anumber)
}
