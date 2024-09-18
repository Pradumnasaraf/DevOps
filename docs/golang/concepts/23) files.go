package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hey I am Pradumna"

	// Create a file
	file, err := os.Create("./hello.txt")

	if err != nil {
		panic(err)
	}

	// Writing content in a file
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length:", length)

	// Reading a file
	data, err := os.ReadFile("./hello.txt")

	if err != nil {
		panic(err)

	}

	// Print the read data
	fmt.Println(string(data))

	//Removing a file
	err = os.Remove("./hello.txt")

	if err != nil {
		panic(err)

	}
}
