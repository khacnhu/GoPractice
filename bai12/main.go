package main

import "fmt"

// Func
func main() {
	var kq int = sum(2, 3)
	fmt.Println("kq = ", kq)

	res := sum(3, 4)
	fmt.Println("RESNEW = ", res)

}

func sum(num1 int, num2 int) int {
	return (num1 + num2)
}
