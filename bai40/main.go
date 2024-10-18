package main

import (
	"fmt"
	"sync"
)

var count = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	count = count + 1

	m.Unlock()

}

// args when run ( go run main.go google.com facebook.com gmail.com )
func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}

	wg.Wait()

	fmt.Println("count =", count)

}
