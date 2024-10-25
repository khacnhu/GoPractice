package main

import (
	"fmt"
)

// hello world
func main() {

	// error when run => bad practice
	ch := make(chan int)
	ch <- 42
	value := <-ch
	fmt.Println("in value = ", value)

	// good practice when run
	chnew := make(chan int)
	go func() {
		chnew <- 42
	}()

	valuenew := <-chnew
	fmt.Println("in value = ", valuenew)

}
