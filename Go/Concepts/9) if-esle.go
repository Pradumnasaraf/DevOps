package main

import (
	"fmt"
)

func main() {
	var name int
	fmt.Println("Enter a number")
	fmt.Scan(&name)

	if name > 10 {
		fmt.Println("Your number is greater than 10")
	} else if name == 10 {
		fmt.Println("Your number is equal to")
	} else {
		fmt.Println("Your number is less than 10")
	}
}

