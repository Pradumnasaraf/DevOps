package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // wait group is used to wait for all go routines to finish

func main() {
	wg.Add(1) // add one to the wait group

	go waitFiveSeconds()
	fmt.Println("Hello from main")

	wg.Wait() // wait for all go routines to finish
}

func waitFiveSeconds() {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello from sayHello")
	wg.Done() // subtract one from the wait group
}
