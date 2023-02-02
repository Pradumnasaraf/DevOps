package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello, playground!"

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the file is: ", length)
	defer file.Close()

	readFile("./test.txt")

}

func readFile(filename string){

	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside a file is \n", string(databyte))
}
