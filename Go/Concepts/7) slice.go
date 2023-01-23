package main

import "fmt"

func main() {
	var name []string // or name := []string{} 

	var firstName = "Pradumna"
	var lastName = "Saraf"

	name = append(name, firstName + " " + lastName) // Wll add element to the index 0 
	name = append(name, firstName + " " + lastName) // Wll add element to the index 1 and so on
	name = append(name, firstName + " " + lastName)

	fmt.Println(len(name))
}
