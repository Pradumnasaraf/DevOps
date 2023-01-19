package main

import (
	"fmt"
)

func main() {

	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
