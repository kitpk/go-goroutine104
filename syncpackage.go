package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mutex
var m sync.Mutex
var count = 10

func doMutex() {
	fmt.Println("FIRST")
	go doWork()
	fmt.Println("SECOND")
	doWork()
	fmt.Println("THIRD")
	time.Sleep(3 * time.Second)
	fmt.Println("DONE")
}
func doWork() {
	m.Lock()
	fmt.Println("LOCK")
	count++
	fmt.Println(count)
	time.Sleep(1 * time.Second)
	m.Unlock()
	fmt.Println("UNLOCK")
}

// WaitGroup
func doWaitGroup() {
	var wg sync.WaitGroup
	job := 4
	for i := 1; i <= job; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work by sleeping
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done\n", id)
}

// Once
func doOnce() {
	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		fmt.Println("Initializing only once")
	}

	doWork := func(workerId int) {
		defer wg.Done()
		fmt.Printf("Worker %d started\n", workerId)
		once.Do(initialize)
		fmt.Printf("Worker %d done\n", workerId)
	}

	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

// Cond
func doCond() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	ready := false

	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()
		for !ready {
			cond.Wait() // Wait for the condition
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()
	}()

	time.Sleep(2 * time.Second)

	mutex.Lock()
	ready = true
	cond.Signal()
	mutex.Unlock()
	fmt.Println("Push signal")

	time.Sleep(1 * time.Second)
	fmt.Println("Work is done.")
}
