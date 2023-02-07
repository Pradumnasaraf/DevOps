package main

import "fmt"

func main() {
	a := 2

	fmt.Println(a)
	defer fmt.Println(a)
	a++
	fmt.Println(a)


}

