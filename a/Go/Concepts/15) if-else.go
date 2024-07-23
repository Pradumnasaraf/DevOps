package main

import (
	"fmt"
)

func main() {
	var name int = 9

	// The pro

	if name > 10 {
		fmt.Println("Your number is greater than 10")
	} else if name == 10 {
		fmt.Println("Your number is equal to 10")
	} else {
		fmt.Println("Your number is less than 10")
	}

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// The if statement can start with a short statement to execute before the condition

	if num := 9; num < 0 { // ; is used to separate the short statement from the condition
		fmt.Println(num, "is negative")
	} else {
		fmt.Println(num, "is positive")
	}

	// if err != nil { // comma ok syntax
	// 	return err
	// }

}
