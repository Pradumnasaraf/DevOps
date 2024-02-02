package main

import "fmt"

func main() {
	greetUser() // Function call
	fmt.Println(addTwoNumbers(2, 3))

	a, b := returnTwoValues() // Multiple return valeus.
	fmt.Println(a, b)

	c, d := addAndSubtract(10, 5)
	fmt.Println(c, d)

	fmt.Println(proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func greetUser() {
	fmt.Println("Hello, welcome to the app")
}

func addTwoNumbers(a int, b int) int { // Function with Para and retun value
	return a + b
}

func returnTwoValues() (int, string){ // Function with multiple return values. 
	return 12, "Hello"
}

func addAndSubtract(a int, b int) (int, int) { // When we have mmultiple return values, we wrap them in () parenthesis.	
	return a + b, a - b
}

func proAdder(value ...int) int { // Function with variable number of parameters. value is a slice of int
	sum := 0
	for _, v := range value { // value is a slice of int
		sum += v
	}
	return sum
}