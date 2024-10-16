package main

import (
	"fmt"
	"time"
)

// Switch case
func main() {

	var i = 4
	switch i {
	case 0:
		fmt.Println("sunday")
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")

	}

	fmt.Println("----------------------")

	switch i {
	case 0:
		fmt.Println("sunday")
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
		fallthrough

	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")

	}

	fmt.Println("----------------------")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Days of week")
	}

}
