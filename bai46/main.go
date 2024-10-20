package main

import (
	"fmt"
	"time"
)

// Your First Concurrent Golang Program | BASICS GOLANG CONCURRENCY |
func main() {
	defer fmt.Println("hay cho doi nhe")
	fmt.Println("hello world")

	go func() {
		fmt.Println("nhu tran in ra gia tri")
	}()

	time.Sleep(1)

	fmt.Println("end")

}
