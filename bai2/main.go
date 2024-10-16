package main

import "fmt"

// variable and constant
func main() {
	var num1 = 10
	fmt.Println("num 1 = ", num1)

	var num2 int64 = 23
	fmt.Println("num 2 = ", num2)

	firstName := "Dr Vipin"
	fmt.Println("name = ", firstName)

	check := true
	fmt.Println("bool = ", !check)

	const pi = 3.14
	// pi = 23 error when change value of pi, because pi is constant value
	fmt.Println("const value pi = ", pi)

}
