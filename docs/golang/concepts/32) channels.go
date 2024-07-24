package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int, 2) // buffered channel. By default, it is unbuffered
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) // read from channel

		
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// close(myChannel)
		ch <- 5 // write to channel
		ch <- 6 // write to channel
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
