package main

import "fmt"

// slice and array
func main() {
	var num []int = []int{12}
	// num[1] = 2
	// num[2] = 3
	num = append(num, 2)
	fmt.Println("res = ", num)

}
