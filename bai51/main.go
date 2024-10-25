package main

import (
	"fmt"
)

// close in channel dont recieve value
func main() {

	// numberP := runtime.NumCPU()

	// fmt.Println("so = ", numberP)
	go func() {
		fmt.Println("a")
	}()

}
