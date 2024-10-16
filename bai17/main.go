package main

import "fmt"

func swap(num1, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}

// call by reference
func main() {

	n1 := 10
	n2 := 20
	swap(&n1, &n2)
	fmt.Println("check n1 = ", n1)
	fmt.Println("check n2 = ", n2)

}
