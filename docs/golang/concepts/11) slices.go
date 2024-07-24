package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits []string // We don't need to specify the size of the array
	// fruits := []string{} - We can also use this syntax to create a slice
	// fruits := make([]string, 2) - We can also use make function to create a slice

	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Orange")
	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Grapes")

	fmt.Println(fruits)
	fmt.Println(len(fruits)) // length of the array
	fmt.Println(fruits[1:])  // from index 1 to end
	fmt.Println(fruits[:3])  // from index 0 to 3 , 3 is not included.
	fmt.Println(fruits[1:3]) // from index 1 to 3 , 3 is not included.

	fmt.Println(sort.StringsAreSorted(fruits)) // check if the array is sorted

	sort.Strings(fruits) // sort the array
	fmt.Println(fruits)

	fmt.Println(sort.StringsAreSorted(fruits)) 
}
