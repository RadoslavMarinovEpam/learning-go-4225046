package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./fromString.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkError(err)
	length, err := io.WriteString(file, "Hello World")
	fmt.Printf("Wrote to file with %v charct\n", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(file string) {
	data, err := os.ReadFile(file)
	checkError(err)
	fmt.Printf("File content: %s\n", string(data))
}
