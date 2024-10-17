package main

import (
	"fmt"
	"os"
)

// args when run ( go run main.go google.com facebook.com gmail.com )
func main() {
	fmt.Println("check os list ", os.Args)
	for _, v := range os.Args {
		fmt.Println("v = ", v)
	}

}
