package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Using bufio package
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n') // text, error syntax. We ignore the error by using _. 
	fmt.Println(text)

	// Using fmt package
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // Scan takes a pointer to a variable
	fmt.Println("Hello", name)
}
