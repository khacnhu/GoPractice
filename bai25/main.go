package main

import (
	"fmt"
	"time"
)

// goroutine
func main() {
	go printMsg("1")
	go printMsg2("2")
	time.Sleep(time.Second)
}

func printMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg+"msg = ", i)
	}
}

func printMsg2(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg+"msg = ", i)
	}
}
