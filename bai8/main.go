package main

import "fmt"

// array
func main() {
	var data = [5]int{1, 3, 5, 2, 5}

	for i := 0; i < len(data); i++ {
		fmt.Println("res = ", data[i])
	}

	var data2 = [2][3]int{{}, {}}
	fmt.Println("data 2 = ", data2)

	var data3 = [2][3]int{{21, 4, 2}, {3, 6, 3}}
	for i := 0; i < len(data3); i++ {
		for j := 0; j < len(data3[i]); j++ {
			fmt.Print("data[i][j] = ", data3[i][j])
			fmt.Print(", ")
		}
		fmt.Println()
	}

}
