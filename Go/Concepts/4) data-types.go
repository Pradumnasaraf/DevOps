package main

import "fmt"

func main(){
	var userName string // Declaring a variable
	var age int

	userName = "Pradumna" // The go compiler will infer the type of the variable
	age = 25

	fmt.Println("Hello", userName, "you are", age, "years old")

}