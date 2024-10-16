package main

import "fmt"

// hello world
func main() {

	var data = [4]int{1, 2, 4, 4}

	for index, value := range data {
		fmt.Println("index = ", index)
		fmt.Println("value = ", value)
		fmt.Println("data in value = ", data[index])
		fmt.Println("-------------------")

	}

	var sdata = make([]int, 3)
	sdata[0] = 2
	sdata[1] = 3
	sdata[2] = 4

	for _, value := range sdata { // _ : index you can replae "_" to index
		fmt.Println("value = ", value)
	}

}
