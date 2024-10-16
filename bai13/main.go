package main

import "fmt"

// return multiple values
func main() {

	kq1, kq2 := Calculator(2, 3)
	fmt.Println("kq 1 = ", kq1)
	fmt.Println("kq 2 = ", kq2)
}

func Calculator(num1, num2 int) (int, int) {
	return num1 + num2, num1 * 10
}
