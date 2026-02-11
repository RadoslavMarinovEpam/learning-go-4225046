package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber {
	case 1:
		result = "It's Monday"
	case 2:
		result = "It's Tuesday"
	case 3:
		result = "It's Wednesday"
	case 4:
		result = "It's Thursday"
	case 5:
		result = "It's Friday"
	default:
		result = "Weekend day"
	}

	fmt.Println(result)
}
