package main

import "fmt"

type calculator struct {
	num1, num2 int
}

func (c calculator) sum() int {
	return c.num1 + c.num2
}

func (c calculator) mul() int {
	return c.num1 * c.num2
}

// golang method
func main() {
	fmt.Println("HELLO WORLD")

	cc1 := calculator{1, 2}
	a := cc1.sum()
	fmt.Println("a = ", a)

	b := cc1.mul()
	fmt.Println("b = ", b)
}
