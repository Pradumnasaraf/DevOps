package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert a string to an integer using strconv package
	numStr := "42"
	num, _ := strconv.Atoi(numStr) // num, _ syntax used to ignore the error.

	fmt.Printf("num: %d, type: %T\n", num, num)

	// Convert an integer to a string using strconv package
	num = 123
	numStr = strconv.Itoa(num)
	fmt.Printf("numStr: %s, type: %T\n", numStr, numStr)

}
