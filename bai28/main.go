package main

import (
	"fmt"
	"time"
)

// select case for channel
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go msg1(ch1)
	go msg2(ch2)

	for i := 0; i < 5; i++ {
		select {
		case r := <-ch1:
			fmt.Println("CHECK CH1 = ", r)
		case r := <-ch2:
			fmt.Println("CHECK CH2 = ", r)
			// default:
			// 	fmt.Println("DEFAULT OKE FINE")
		}
	}

	// time.Sleep(time.Second * 2)

}

func msg1(ch1 chan string) {
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("Message msg1= %d", i)
		time.Sleep(time.Second)
	}

}

func msg2(ch2 chan string) {
	for i := 0; i < 5; i++ {
		ch2 <- fmt.Sprintf("Message msg2 = %d", i)
		time.Sleep(time.Second)

	}
}
