package main

import "fmt"

func main() {
	var name []string // or name := []string{} 

	var firstName = "Pradumna"
	var lasttName = "Saraf"

	name = append(name, firstName + " " + lasttName) // Wll add elemnt to the index 0 
	name = append(name, firstName + " " + lasttName) // Wll add elemnt to the index 1 and so on
	name = append(name, firstName + " " + lasttName)

	fmt.Println(len(name))
}
