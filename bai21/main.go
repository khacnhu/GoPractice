package main

import "fmt"

// defer
func main() {
	defer fmt.Println("in ra dong 1")
	fmt.Println("in ra dong 2")

	for i := 0; i < 5; i++ {
		defer fmt.Println("i = ", i) // last in first out
	}

}
