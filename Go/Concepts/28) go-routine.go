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

	//runThisInMain()
}

func runThisInMain() {
	for {
		var input string
		go sayHello() // go keyword is used to create a go routine. This is a concurrent execution of the function, meaning it will run in the background and will not block the main thread.
		fmt.Print("Enter Text: ")
		fmt.Scanln(&input)
		fmt.Println("Your Input was:", input)
	}
}

func sayHello() {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello from sayHello")
	
}
