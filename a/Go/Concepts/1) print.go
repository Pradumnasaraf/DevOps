package main

import "fmt"

func main(){
	var name = "Pradumna"
	fmt.Print("Hello World") // Print does not add a new line at the end
	fmt.Println("Hello World") // Println adds a new line at the end
	fmt.Println("Hello World" , name) // , adds a space between the strings
}