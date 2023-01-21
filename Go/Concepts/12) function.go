package main

import "fmt"

func main() {
	greetUser()
	fmt.Println(addTwoNumbers(2, 3))

}

func greetUser() {
	fmt.Println("Hello, welcome to the app")
}

func addTwoNumbers(a int, b int) int { // Function with Para
	return a + b
}