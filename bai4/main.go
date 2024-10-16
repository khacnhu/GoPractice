package main

import "fmt"

// For loop
func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("----------------")

	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("i = ", i)
	}

	fmt.Println("----------------")

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("i = ", i)
	}

}
