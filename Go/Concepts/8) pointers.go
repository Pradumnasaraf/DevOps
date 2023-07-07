package main

func main(){

	// var ptr *int

	myInt := 42
	ptr := &myInt // & is used to get the address of the variable
	// var ptr *int = &myInt

	println("myInt:", ptr) // ptr is the address of the variable
	println("myInt:", *ptr) // * is used to get the value stored at the address

}