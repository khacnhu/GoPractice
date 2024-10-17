package main

import "fmt"

type NumberReturn interface {
	int64 | float64
}

// func generic
func IntFloatSum[K comparable, V NumberReturn](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// generic
func main() {
	ints := map[string]int64{
		"first":  20,
		"second": 30,
	}
	fmt.Println("res = ", IntSum(ints))

	floats := map[string]float64{
		"first":  23.2,
		"second": 12.6,
	}
	fmt.Println("res = ", FloatSum(floats))
	fmt.Println("------------")
	fmt.Println("res = ", IntFloatSum(ints))
	fmt.Println("res = ", IntFloatSum(floats))

}

func IntSum(m map[string]int64) int64 {
	var s int64

	for _, v := range m {
		s += v
	}

	return s

}

func FloatSum(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}
