package main

import "fmt"

// hello world
func main() {
	c1 := make(chan int)
	fmt.Println("push c1:")
	c1 <- 10
	g1 := <-c1
	fmt.Println("get g1: ", g1)

}
