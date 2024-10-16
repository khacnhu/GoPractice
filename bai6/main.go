package main

import "fmt"

// switch as if
func main() {
	var i = 4
	switch {
	case i >= 0:
		fmt.Println("sunday")
	case i >= 1:
		fmt.Println("monday")
	case i >= 2:
		fmt.Println("tuesday")
	case i >= 3:
		fmt.Println("wednesday")
	case i >= 4:
		fmt.Println("thursday")
	case i >= 5:
		fmt.Println("friday")
	case i >= 6:
		fmt.Println("saturday")

	}
}
