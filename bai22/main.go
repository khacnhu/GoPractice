package main

import (
	"fmt"
)

type devidebyzero struct {
	textError string
}

func (c devidebyzero) Error() string {
	return c.textError
}

func errorDevideByZero() error {
	return devidebyzero{"Devide by zero will be failed in math"}
}

// exception handling
func main() {
	kq, err := getDivision(10, 5)
	fmt.Println("kq = ", kq)
	fmt.Println("err = ", err)
	fmt.Println("----------")

	kqnew, errnew := getDivision(10, 0)
	fmt.Println("kq = ", kqnew)
	fmt.Println("err = ", errnew)
	fmt.Println("----------")

}

func getDivision(num1, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errorDevideByZero()
	}
	return num1 / num2, nil
}
