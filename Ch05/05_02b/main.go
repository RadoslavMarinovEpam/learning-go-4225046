package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	fmt.Println(dog.Speak3Times())

}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

func (d Dog) Speak3Times() string {
	return fmt.Sprintf("The %v says %v! %v! %v!", d.Breed, d.Sound, d.Sound, d.Sound)
}
