package main

// hello world
func main() {

}

type Node struct {
	value int
	pre   *Node
	next  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoubleLinkedList) InsertAtEnd(value int) {
	newNode := &Node{}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode = list.tail
		list.tail = newNode
	}

}

func (list *DoubleLinkedList) InsertAtBegin(value int) {
	newNode := &Node{}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.pre = newNode
		list.head = newNode
	}

}
