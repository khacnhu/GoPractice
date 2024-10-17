package main

import "fmt"

type Student struct {
	name   string
	course string
	age    int64
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) GetCourse() string {
	return s.course
}

func (s *Student) SetCourse(course string) {
	s.course = course
}

func (s *Student) GetAge() int64 {
	return s.age
}

func (s *Student) SetAge(age int64) {
	s.age = age
}

// escapsulation as class
func main() {
	s := Student{}
	s.SetName("nhutran")
	s.SetCourse("python")
	s.SetAge(12)
	fmt.Println("s = ", s)

}
