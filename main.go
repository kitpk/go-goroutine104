package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine
	go doSomething1()
	doSomething2()
	// wait doSomething1 and doSomething2 work done
	time.Sleep(time.Second)
	fmt.Println("Done")

	// Channel
	doChannel1()
	doChannel2()
	doChannelLoopReceive()
	doChannelLoopSend()
	doChannelSelect()

	// Sync packagea
	doMutex()
	doWaitGroup()
	exampleMutexAndWaitGroup()
	doOnce()
	doCond()
}

func doSomething1() {
	fmt.Println("Doing something 1")
}

func doSomething2() {
	fmt.Println("Doing something 2")
}
