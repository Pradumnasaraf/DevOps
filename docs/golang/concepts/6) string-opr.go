package main

import (
	"fmt"
	"strings"
)

func main() {

	var name = "John"
	fmt.Println(len(name))               // It will print the length of the string
	fmt.Println(strings.ToUpper(name))   // It will print the string in uppercase
	fmt.Println(strings.ToLower(name))   // It will print the string in lowercase
	fmt.Println(strings.Title(name))     // It will print the string in title case
	fmt.Println(strings.Repeat(name, 3)) // It will repeat the string 3 times
	fmt.Println(strings.Replace(name, "J", "K", 1)) // It will replace the first occurrence of J with K
	fmt.Println(strings.Split(name, "")) // It will split the string into a slice of characters
}
