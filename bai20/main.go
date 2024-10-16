package main

import "fmt"

// interface
func main() {
	fmt.Println("HELLO WORLD")

	cc := calculator{2, 3}
	useCalculator(cc)

}

type ICalculator interface {
	sum() int
	mul() int
}

type calculator struct {
	num1, num2 int
}

func (c calculator) sum() int {
	return c.num1 + c.num2
}

func (c calculator) mul() int {
	return c.num1 * c.num2
}

func useCalculator(cc ICalculator) {
	fmt.Println("tong = ", cc.sum())
	fmt.Println("tich = ", cc.mul())
}
