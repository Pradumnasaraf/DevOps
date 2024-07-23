package main

import (
	"fmt"
	"sync"
)

func main() {
	// When channel is closed, the value read from the channel is the zero value and at the same time maybe the write value can also be zero value.
	// So we need to check if the channel is closed or not by using the second return value of the read operation.

	myChannel := make(chan int, 2) 
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch 
		fmt.Println(val, isChannelOpen)		
		wg.Done()
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch <- 5          uncomment this line to see the difference
		close(myChannel)
		wg.Done()

	}(myChannel, wg)

	wg.Wait()
}
