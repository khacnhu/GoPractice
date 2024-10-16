package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// anonymous function
func main() {
	sum := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println("kq = ", sum(3, 5))
	fmt.Println("--------------------")
	count := counter()
	fmt.Println("check index create ", count())
	fmt.Println("check index create ", count())

	fmt.Println("check index create ", count())

}
