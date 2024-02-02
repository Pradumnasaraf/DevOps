package main

import "fmt"

func main() {
	// remove an element from a slice
	var cars = []string{"BMW", "Audi", "Mercedes", "Ford", "Fiat"}

	fmt.Println(cars)
	var index int = 2
	cars = append(cars[:index], cars[index+1:]...) // remove the element at index 2. :index is from 0 to index-1, [index+1:] is from index+1 to end. We can append both elements as well as a slice. ... is to denote a slice
	fmt.Println(cars)

}
