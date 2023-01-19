package main

import "fmt"

func main(){
	var temp = 20 // tempparature can change 
	const myCity = "Mumbai" // name of the city is constant
	name := "Pradumna" // := is a shortcut for declaring and initializing a variable

	fmt.Println("The temperature in", myCity, "is", temp, "degrees")
	fmt.Println("My name is", name)
}