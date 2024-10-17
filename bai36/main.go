package main

import "fmt"

type Student struct {
	name string
	age  int64
}

func (s *Student) GetInfoStd() {
	fmt.Printf("information name = %s , age = %d ", s.name, s.age)
}

// structure as class
func main() {
	s := Student{name: "nhu tran", age: 12}
	s.GetInfoStd()
}
