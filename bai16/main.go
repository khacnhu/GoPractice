package main

import "fmt"

// Pointer
func main() {
	var num1 int = 20
	fmt.Println("num1 = ", num1)
	fmt.Println("check pointer num 1 = ", &num1)

	var num2 = &num1
	fmt.Println("chek num 2 = ", num2)
	*num2 = 23
	fmt.Println("again num 1 = ", num1)
	fmt.Println("again num 2 = ", *num2)
	fmt.Println("check pointer of num2 = ", &num2)
}
