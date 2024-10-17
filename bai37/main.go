package main

import "fmt"

type Student struct {
	firstname string
	lastname  string
	age       int
}

func (s *Student) GetFullName() string {
	return fmt.Sprintf("information name = %s %s ", s.firstname, s.lastname)
}

type topic struct {
	title   string
	content string
	student Student
}

func (p topic) details() {
	fmt.Println("title TOPIC: ", p.title)
	fmt.Println("Content TOPIC: ", p.content)
	fmt.Println("Student TOPIC: ", p.student.GetFullName())

}

type website struct {
	topics []topic
}

func (w *website) contents() {
	fmt.Println("start content of website")
	for _, value := range w.topics {
		value.details()
		fmt.Println("-------------")
	}
}

// structure as class
func main() {
	s1 := Student{firstname: "nhu tran", lastname: "tran", age: 12}
	s2 := Student{firstname: "le trang", lastname: "thi", age: 43}

	t1 := topic{
		title:   "title one",
		content: "noi dung ben trong la nhu the nao vay moi nguio",
		student: s1,
	}

	t2 := topic{
		title:   "two tilte",
		content: "content is written by foreigners",
		student: s1,
	}

	t3 := topic{
		title:   "third tilte",
		content: "AAAAAAAAAAAAAAAAA",
		student: s2,
	}

	web := website{
		topics: []topic{t1, t2, t3},
	}

	web.contents()

}
