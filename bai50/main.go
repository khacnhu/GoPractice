package main

import "fmt"

// close in channel dont recieve value
func main() {
	queue := make(chan int)

	close(queue)

	go func() {
		queue <- 2

	}()

	v := <-queue
	fmt.Println("v = ", v)

}
