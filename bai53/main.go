package main

import "fmt"

func main() {
	a := 100
	var pointer *int

	pointer = &a

	fmt.Println("valuea address pointer = ", pointer)
	fmt.Println("value of pointer = ", *pointer)

	fmt.Println("-------------------------------------")
	// cach so 2
	b := 200
	p2 := new(int)

	p2 = &b
	fmt.Println("value address = ", p2)
	fmt.Println("value = ", *p2)

	fmt.Println("-------------------------------------")
	var pointer3 *int
	fmt.Println("pointer = nil ", pointer3)

	fmt.Println()
	// no giong voi reference type
	var pointer4 *int
	c := 20
	pointer4 = &c

	*pointer4 = 42

	fmt.Println("c = ", c)
	fmt.Println("value pointer 4 = ", *pointer4)

	// pointer in array in golang
	fmt.Println()
	arr := [3]int{1, 2, 3}

	pointer5 := &arr
	fmt.Println("array address ", pointer5)
	pointer5[2] = 23
	fmt.Println("value of arr pointer ", *pointer5)
	fmt.Println("value of arr new = ", arr)
	fmt.Printf("value T = %T ", pointer5)
}
