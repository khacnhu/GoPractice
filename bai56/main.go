package main

import "fmt"

type Animal interface {
	speak()
}

type Movement interface {
	move()
}

type NextAnimal interface {
	Movement
	Animal
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("woaaaaa")
}

func (d Dog) move() {
	fmt.Println("runnnnnnnn moveeeeeeeeee")
}

// interface
func main() {

	//-------------- cach 1 ---------------- interface ---------
	// var animal Animal
	// animal = Dog{}
	// animal.speak()

	//-----------------cach 2----------------- multiple interface -----------------
	dog := Dog{}

	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	fmt.Println()
	// ------------- cach 3------------------- nested interface -------------
	dogNew := Dog{}

	var ma NextAnimal = dogNew
	ma.move()
	ma.speak()

}
