package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Routine")
		mut.Lock() // Lock the memory location of score
		score = append(score, 1)
		mut.Unlock() // Unlock the memory location of score
		wg.Done()
	}(wg, mut) // Pass the pointer of WaitGroup and Mutex
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
