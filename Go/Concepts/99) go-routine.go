package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		var input string
		go sayHello() // go keyword is used to create a go routine
		fmt.Print("Enter Text: ")
		fmt.Scanln(&input)
		fmt.Println("Your Input was:", input)
	}
}

func sayHello() {
	time.Sleep( 5 * time.Second)
	fmt.Println("Hello from sayHello")
}
