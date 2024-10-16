package main

import (
	"fmt"
	"sync"
	"time"
)

// hello world
func main() {

	// ch := make(chan string, 2)

	// // ch <- "this is value 1 in channel"
	// // ch <- "do you understand"

	// // close(ch)
	// // for chv := range ch {
	// // 	fmt.Println("chv = ", chv)
	// // }

	chNew := make(chan string, 1)
	go displayMessageWait(chNew)
	for chv := range chNew {
		fmt.Println("chv = ", chv)
	}
}

func displayMessage(ch chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 2)
		ch <- fmt.Sprintf("this is value %d n channel ", i)
	}
	close(ch)
}

func displayMessageWait(ch chan string) {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 10)
			ch <- fmt.Sprintf("this is value %d n channel ", i)
		}()

	}

	wg.Wait()
	close(ch)
}
