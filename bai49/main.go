package main

import "fmt"

// create custom struct tags in golang is awesome!
func main() {
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}

		done <- true

	}()
	for {

		select {
		case v := <-queue:
			fmt.Println("res = ", v)
		case q := <-done:
			fmt.Println("q = ", q)
			return
		default:
			fmt.Println("khong co ket qua")
		}
	}

}
