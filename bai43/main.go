package main

import "fmt"

// queue in golang
func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)

	queue.Dequeue()
	fmt.Println("dequeue = ", queue)

	val, ok := queue.Peek()
	if !ok {
		fmt.Println("empty")
	} else {
		fmt.Println("list ", val)

	}

}

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// return the first in element but remove it in list
func (q *Queue) Dequeue() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false // return false if the queue is empty
	}

	frontElement := q.elements[0]

	q.elements = q.elements[1:]

	return frontElement, true

}

// peek: return the first element without remove it
func (q *Queue) Peek() (int, bool) {
	if len(q.elements) == 0 {
		return 0, false
	}

	return q.elements[0], true
}
