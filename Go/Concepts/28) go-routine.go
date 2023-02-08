package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello")
	greeter("World")

}

func greeter(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func runThisInMain() {
	for {
		var input string
		go sayHello() // go keyword is used to create a go routine. It will not wait for the function to finish.
		fmt.Print("Enter Text: ")
		fmt.Scanln(&input)
		fmt.Println("Your Input was:", input)
	}
}

func sayHello() {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello from sayHello")
}
