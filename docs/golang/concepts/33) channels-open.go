package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	myCh := make(chan int, 1)

	wg.Add(2)
	go func(channel chan int, wait *sync.WaitGroup) {
		fmt.Println("First Go Routine")
		wait.Done()

		value, isChannelOpen := <-myCh

		if !isChannelOpen {
			fmt.Println("Channel Closed")
		}
		fmt.Println(value)
	}(myCh, wg)

	go func(channel chan int, wait *sync.WaitGroup) {
		fmt.Println("Second Go Routine")
		// myCh <- 3
		close(myCh)
		wait.Done()
		// myCh <- 5
	}(myCh, wg)

	wg.Wait()
}
