package main

import (
	"fmt"
	"math/rand"
	"time"
)

// hello world
func main() {
	result := make(chan string)
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second)
		go getSum(result)
		go getResult(result)
	}
	// close(result)

}

func getRand() int {
	return rand.Intn(100)

}

func getSum(result chan string) {
	n1 := getRand()
	n2 := getRand()
	sum := n1 + n2
	result <- fmt.Sprintf("The sum of %d and %d = %d ", n1, n2, sum)
	time.Sleep(time.Second * 2)

}

func getResult(result chan string) {
	fmt.Println("result = ", <-result)
}
