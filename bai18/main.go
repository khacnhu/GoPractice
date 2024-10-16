package main

import "fmt"

// struct
func main() {
	type person struct {
		name          string
		qualification string
	}

	p1 := person{"dr. trankhacnu", "phd dont study uni"}

	fmt.Println("person = ", p1)

	fmt.Println("name = ", p1.name)
	fmt.Println("qualification = ", p1.qualification)

	sp1 := &p1
	sp1.name = "dai ca giang ho"
	fmt.Println("sp1 name = ", sp1.name)
	fmt.Println("again p1 name = ", p1.name)
}
