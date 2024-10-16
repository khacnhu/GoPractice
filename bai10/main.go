package main

import "fmt"

// map
func main() {
	var temp = make(map[string]int)
	temp["nhutran"] = 34
	temp["anaconda"] = 12
	temp["sitcom"] = 23

	fmt.Println("res = ", temp)
	fmt.Println("nhutran res = ", temp["nhutran"])

	res, check := temp["nhutran"]
	fmt.Println("var res = ", res)
	fmt.Println("var check = ", check)

	resnew, checknew := temp["baba"]
	fmt.Println("var resnew = ", resnew)
	fmt.Println("var checknew = ", checknew)
}
