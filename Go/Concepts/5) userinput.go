package main

import "fmt"

func main() {
	var name string

	fmt.Scan(&name) // Scan takes a pointer to a variable

	fmt.Println("Hello", name)
}