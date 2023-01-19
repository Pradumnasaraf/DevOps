## Go

Go is statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

In Go, everything is a package. A package is a collection of source files in the same directory that are compiled together. A package can be imported by other packages. `main` is a special package that defines a standalone executable program, not a library. 

### Installation

- [Download](https://golang.org/dl/) and install Go

### Resources

- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/welcome/1)
- [Go Playground](https://play.golang.org/)
- [Go Documentation](https://golang.org/doc/)
- [Codeacademy Free course](https://www.codecademy.com/learn/learn-go)

### Features

- Garbage collected
- Multithreading
- concurrency

### Packages

- fmt - formatted I/O with functions analogous to C's printf and scanf.
- os - provides a platform-independent interface to operating system functionality.
- strconv - implements conversions to and from string representations of basic data types.
- time - provides functionality for measuring and displaying time.
- math - provides basic constants and mathematical functions.
- net/http - provides HTTP client and server implementations.
- encoding/json - implements encoding and decoding of JSON.

### Hello World - Running a program

```go
package main // package declaration

import "fmt" // import fmt package

func main() {
    fmt.Println("Hello, World!")
}
```

We can run and compile the program using the following command:

```bash
$ go run hello.go
```

### Print 

Placeholders for `printf` : 

- %v - value in default format
- %T - type of value
- %t - boolean
- %d - decimal integer
- %b - binary integer
- %c - character
- %x - hexadecimal integer
- %f - floating point number

Escape sequences:

- \n - newline
- \t - tab

> Examples in `Concepts/1) print.go`

### Variables

- `var` - variables with initializers.
- `const` - declares a constant value. Can't be updated.
- `:=` - short variable declaration. Can be used only inside a function. Called Syntax sugar.


> Examples in `Concepts/2) variables.go`

### Data Types

When we declare with a value we don't need to specify the type. Go will infer the type from the value. But when we declare without a value we need to specify the type.

Eg: 

```go
var name string

name = "John"
```

- `bool` - it can be either true or false. Example: `true`
- `string` - string of characters. Example: `"Hello, World!"`
- `int` - integer number. Example: `42`
- `float64` - floating point number. Example: `3.14`
- `uint` - unsigned integer. Example: `42`
- `byte` - alias for uint8. Example: `42`

### Scan

`fmt.Scan` reads text from standard input, scanning the text read into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

```go
fmt.Scan(&name)
```

### Pointers

A pointer holds the memory address of a value.

```go
var name = "John"
fmt.Println(&name) // it will print the memory address of the variable name
```


