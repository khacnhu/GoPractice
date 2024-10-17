package main

import (
	"fmt"
	"os"
)

// escapsulation as class
func main() {
	fmt.Println("check os list ", os.Args)
	for _, v := range os.Args {
		fmt.Println("v = ", v)
	}

}
