package main

import "fmt"

// create custom struct tags in golang is awesome!
func main() {

	a := make(chan string, 2)

	a <- "alo"
	a <- "pdf"

	fmt.Println("a = ", <-a)
	fmt.Println("a2 = ", <-a)

	close(a)

	// print lan thu 3 bi deadlock if dont close above
	fmt.Println("a3 = ", <-a)

}
