package hello

import "fmt"

func Hello() { // Function name starts with a capital letter so it is exported. It can be called from outside the package
	fmt.Println("Hello from hello.go")
}