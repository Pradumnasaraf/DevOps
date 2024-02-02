package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Process to create, write and read a file

	content := "Hey I am Pradumna"

	//"." means current directory
	file, err := os.Create("./hello.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length:", length)

	data, err := os.ReadFile("./hello.txt")

	if err != nil {
		panic(err)

	}

	// data is of type []byte so we need to convert it to string
	fmt.Println(string(data))
}
