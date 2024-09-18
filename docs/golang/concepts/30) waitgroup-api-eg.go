package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup 
var mut sync.Mutex   
var statusCodes = []int{}

func main() {

	endpoints := []string{
		"https://www.google.com",
		"https://www.pradumnasaraf.dev",
		"https://www.go.dev",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.medium.com",
	}

	for _, endpoint := range endpoints {
		go checkStausCode(endpoint)
		wg.Add(1)
	} 

	wg.Wait()
	fmt.Print(statusCodes)
}

func checkStausCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Problem in the endpoint", endpoint)
	} else {
		mut.Lock()
		statusCodes = append(statusCodes, res.StatusCode)
		mut.Unlock()
		fmt.Printf("Status code for %s is %d\n", endpoint, res.StatusCode)
	}

}
