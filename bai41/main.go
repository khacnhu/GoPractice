package main

import "fmt"

// Linked List algorthm
func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Insert(50)
	list.Insert(60)

	list.Display()
	list.Delete(20)
	list.Display()

}

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(value int) {
	newNode := &Node{value: value, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
			fmt.Println("current next la cai gi: ", current)
		}
		current.next = newNode
	}

}

func (list *LinkedList) Display() {
	current := list.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}

	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")

}

func (list *LinkedList) Delete(value int) {
	current := list.head

	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	if list.head.value == value {
		list.head = current.next
		return
	}

	for current.next != nil && current.next.value != value {
		current = current.next
	}

	current.next = current.next.next

}
