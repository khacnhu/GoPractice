package main

import (
	"fmt"
	"time"
)

// timer
func main() {

	// timer1 := time.NewTimer(10 * time.Second)
	// fmt.Println("before timer start")
	// <-timer1.C
	// fmt.Println("After 2 second timer expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Inside function timer2 expired")
	}()
	stop := timer2.Stop()

	time.Sleep(time.Second * 3)

	if stop {
		fmt.Println("timer stopped before expired")
	}

}
