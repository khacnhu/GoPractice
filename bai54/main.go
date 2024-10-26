package main

import "fmt"

type NoName struct {
	string
	int
}

type StudentInfo struct {
	class string
	email string
	age   int
}

type Student struct {
	id   int
	name string
	info StudentInfo
}

// struct
func main() {
	a := NoName{
		"con chim",
		23,
	}

	fmt.Println("a = ", a)

	b := NoName{
		"anaconda",
		43,
	}

	fmt.Println("b = ", b)

	fmt.Println()

	studentNested := Student{
		id:   2,
		name: "nhutran",
		info: StudentInfo{
			class: "ly",
			email: "tkn@gmail.com",
			age:   23,
		},
	}
	fmt.Println("nested = ", studentNested)

	var studentOne Student
	fmt.Println("student one = ", studentOne)

}
