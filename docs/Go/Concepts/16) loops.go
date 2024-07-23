package main

import (
	"fmt"
)

func main() {

	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	sum := 0
	for i := 0; i < 10; i++ { //for [initialization]; [condition]; [post] { }
		sum += i
	}
	fmt.Println(sum)

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := 0; i < len(days); i++ { //	Simple wy to iterate over an array
		fmt.Println(days[i])
	}

	// for [index] := range [array] { }
	for i := range days { // Better way to iterate over an array
		fmt.Println(days[i])
	}

	// for [index], [value] := range [array] { }
	for i, day := range days { // This is the preferred way to iterate over an array
		fmt.Println(i, day) // i is the index and day is the value
	}

	// for _, [value] := range [array] { }
	for _, day := range days { // we can ignore the index by using _
		fmt.Println(day)
	}

}
