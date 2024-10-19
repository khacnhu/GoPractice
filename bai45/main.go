package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	per() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type equiTriangle struct {
	length float64
}

func (e equiTriangle) area() float64 {
	return e.length * 10000
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) per() float64 {
	return 4 * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) per() float64 {
	return 2 * math.Pi * 3.14
}

// interface practice more
func main() {

	// square := square{length: 2}
	// circle := circle{radius: 3}

	listshape := []shape{square{length: 2}, circle{radius: 3}}

	for _, val := range listshape {
		fmt.Println("are = ", val.area())
		fmt.Println("per = ", val.per())
		fmt.Println("-----------------")
	}

}
