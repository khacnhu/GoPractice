package main

import "fmt"

// type String string

// func (c String) append(str string) string {
// 	return fmt.Sprintf("%s %s ", c, str)
// }

type Student struct {
	id   int
	name string
}

// Method with value reciever cann not change name
func (s Student) changeName(namenew string) {
	s.name = namenew
}

// Method with pointer receiver can change name
func (s *Student) change(namenew string) {
	s.name = namenew
}

// method funcion
func main() {

	// cnew := String("a")

	// val := cnew.append("tran khac nhu")
	// fmt.Println("val = ", val)

	student := &Student{id: 1, name: "tran khac nhu"}

	student.change("le trang thi thanh")
	fmt.Println("student pointer = ", student)

	studentNew := Student{id: 2, name: "bao bach tran"}

	studentNew.changeName("kng dom boy")
	fmt.Println("studnet with value reciever ", studentNew)
}
