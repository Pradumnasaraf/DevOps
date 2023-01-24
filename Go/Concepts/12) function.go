package main

import "fmt"

func main() {
	greetUser()
	fmt.Println(addTwoNumbers(2, 3))

	a, b := returnTwoValues()
	fmt.Println(a, b)

	c, d := addAndSubtract(10, 5)
	fmt.Println(c, d)
}

func greetUser() {
	fmt.Println("Hello, welcome to the app")
}

func addTwoNumbers(a int, b int) int { // Function with Para
	return a + b
}

func returnTwoValues() (int, string){ // Function with multiple return values. 
	return 12, "Hello"
}

func addAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}
