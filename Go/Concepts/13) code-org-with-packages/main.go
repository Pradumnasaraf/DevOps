package main

import (
	"code-org/name"
	"code-org/hello"

)

func main() {
	name.Name() //Name() functtion is exported so we can call it from outside the package
	hello.Hello() //Hello() functtion is exported so we can call it from outside the package
	wish() // Wish exists in the same package as main so we can call it directly without importing it 	
}
