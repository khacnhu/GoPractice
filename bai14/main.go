package main

import "fmt"

// variadic function
func main() {
	var kq int = sum(1, 2, 23)
	fmt.Println("res = ", kq)

}

func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		s = s + num
	}
	return s
}
