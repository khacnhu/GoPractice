package main

import (
	"fmt"
)

// defer panic recovery again
func main() {
	// file, _ := os.Create("file.txt")
	// _, _ = io.WriteString(file, "hello tran khac nhu")

	// file.Close()
	defer handlePanic()
	fmt.Println("start")
	panic("error exception")

}

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("recover ok")
	}
}
