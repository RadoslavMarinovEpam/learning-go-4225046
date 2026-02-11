package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello")

	go func(message string) {
		fmt.Println(message)
	}("Hello from anonymous function")

	fmt.Println("helo main")
	time.Sleep(time.Second * 2)
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
