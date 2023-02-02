package main

import "fmt"

func main() {
	var nameArray1 = [5]string{}
	var nameArray [5]string

	nameArray[0] = "John"
	nameArray1[0] = "Doe"
	nameArray1[4] = "Doe"
	//nameArray1[14] = "Doe"  - Not possible as array size is 5


	fmt.Println(nameArray)
	fmt.Println(nameArray1)
	fmt.Println(nameArray1[0])
	fmt.Printf("Array type: %T\n" , nameArray1)
	fmt.Println(len(nameArray1))

}
