package main

import (
	"fmt"
	"sync"
)

// Example Using Mutex and WaitGroup
type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.value
}

func exampleMutexAndWaitGroup() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value())
}
