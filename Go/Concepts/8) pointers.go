package main

func main(){

	// var ptr *int

	myInt := 42
	ptr := &myInt

	println("myInt:", ptr) // ptr is the address of the variable
	println("myInt:", *ptr) // *ptr is the value of the pointer



	
}