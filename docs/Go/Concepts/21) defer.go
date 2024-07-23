package main

import "fmt"

func main() {
	defer fmt.Println("This is the first defer statement")
	defer fmt.Println("This is the second defer statement")
	defer fmt.Println("This is the third defer statement")
	fmt.Println("Hello, playground")

	myDefer( )

}

func myDefer(){
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}