package main

import (
	"fmt"
	"time"
)

func doChannel1() {
	fmt.Println("Doing Channel 1")
	ch := make(chan int, 1)
	ch <- 1

	v := <-ch
	fmt.Println(v)
}

func doChannel2() {
	fmt.Println("Doing Channel 2")
	ch := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Call 1")
		ch <- 2
		ch <- 3
	}()
	fmt.Println("Call 2")

	v := <-ch
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
}

func doChannelLoopReceive() {
	fmt.Println("Doing Channel Loop Receive")
	ch := make(chan int)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	ch <- 10
	ch <- 20
	ch <- 30
}

func doChannelLoopSend() {
	fmt.Println("Doing Channel Loop Send")
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func doChannelSelect() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		channel1 <- 100
		close(channel1)
	}()

	go func() {
		channel2 <- 200
		close(channel2)
	}()

	closedChannel1, closedChannel2 := false, false

	for {
		if closedChannel1 && closedChannel2 {
			break
		}

		select {
		case v, ok := <-channel1:
			if !ok {
				closedChannel1 = true
				continue
			}
			fmt.Println("Channel 1", v)
		case v, ok := <-channel2:
			if !ok {
				closedChannel2 = true
				continue
			}
			fmt.Println("Channel 2", v)
		}
	}
}
