package main

import "fmt"

func main() {
	var num int = 42
	var numFloat float64 = float64(num)

	var numFloat2 = 3.14
	var numInt = int(numFloat2)

	fmt.Println(numFloat)
	fmt.Println(numInt)
}

