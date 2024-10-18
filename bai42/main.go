package main

import "fmt"

// STACK IN GOLANG
func main() {
	stack := Stack{}

	// Push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	// Peek at the top element
	top, _ := stack.Peek()
	fmt.Printf("Top element: %d\n", top)

	// Pop elements from the stack
	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Printf("Popped: %d\n", value)
	}

	// Try popping from an empty stack
	if value, ok := stack.Pop(); !ok {
		fmt.Println("Stack is empty, cannot pop")
	} else {
		fmt.Println("value ", value)
	}
}

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false // return false if stack is empty
	}

	topElement := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]
	return topElement, true

}

func (s *Stack) Peek() (int, bool) {
	if len(s.elements) == 0 {
		return 0, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0

}
