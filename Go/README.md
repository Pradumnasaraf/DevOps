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
- [Golang Tutorial TechWorld with Nana](https://youtu.be/yyUHQIec83I)

### Features

- Garbage collected
- Multithreading
- concurrency

### Packages

Module is a collection of related Go packages that are versioned together as a single unit.

- fmt - formatted I/O with functions analogous to C's printf and scanf.
- os - provides a platform-independent interface to operating system functionality.
- strconv - implements conversions to and from string representations of basic data types.
- time - provides functionality for measuring and displaying time.
- math - provides basic constants and mathematical functions.
- net/http - provides HTTP client and server implementations.
- encoding/json - implements encoding and decoding of JSON.

#### Strings

- strings.Contains(s, substr) - checks whether a substring exists within a string.
- strings.Fields(s) - splits a string into an array of substrings separated by spaces.

### Hello World - Running a program

```go
package main // package declaration

import "fmt" // import fmt package

import ( "fmt" "os") // import multiple packages

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

### Arrays

An array is a numbered sequence of elements of a single type with a fixed length.

We declare an array as follows:

```go
var arr [5]int  // array of 5 integers
```

### Slice

A slice is a segment of an array. It is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

We declare a slice as follows:

```go
var names []int  // slice of integers

names = append(names, 1) // append is a built-in function to add an element to the slice

```

### Loops

In Go, there is only one looping construct, the `for` loop.

```go
for i := 0; i < 5; i++ {
fmt.Println(i)
}
```

- `range` - The range form of the for loop iterates over a slice or map.

```go
for i, name := range names {
fmt.Println(i, name)
}
```

- `_` - The blank identifier is a special identifier that is used to ignore values when multiple values are returned from a function. It is used to ignore the index in the above example.

```go
for _, name := range names {
fmt.Println(name)
}
```

### if else

```go
if num > 0 { // this condition will only be true if num is greater than 0
fmt.Println("Positive")
} else if num == 0 { // this condition will only be true if num is equal to 0
fmt.Println("Equal to zero")
} else {   // this condition will only be true if num is less than 0
fmt.Println("Negative")
}
```

We can check also write `data == false` or `!data`

#### Break and Continue

- `break` - The break statement terminates the loop or switch statement and transfers execution to the statement immediately following the loop or switch.

```go

for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}
```

- `continue` - The continue statement terminates the current iteration of the loop, and resumes execution at the next iteration. It can be used only within an iterative or switch statement and only within the body of that statement.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        continue
    }
    fmt.Println(i)
}
```

### Switch

The switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

```go
switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

### Functions

- A function is a block of code that performs a specific task. It is a reusable piece of code.

```go
func add(x int, y int) int { // We can specify the type of the parameters
    return x + y
}
```

- A function can return multiple values.

```go
func swap(x, y string) (string, string) { // When we have more than one return value we put them in parenthesis ().
return y, x
}

a, b := swap("hello", "world") // We can get the return values using multiple assignment
```

### Package level variables

- We can declare variables at the package level. They are called package level variables.

NOTE: We can't use the short variable declaration operator `:=` to declare package level variables.

```go
var name string = "John" // It will be available to all the functions in the package

func main() {
    fmt.Println(name)
}
```

### Code Organization

We can organize our code by putting functions/ variables in different files and can use them in the main file or calling the main function from other files.

Also, we can have multiple packages in a single directory.

### Expoting and Importing

- We can export a function/ variable by capitalizing the first letter of the function/ variable name. Now we can use it in other packages.

```go
func Add(x, y int) int {
    return x + y
}

var Name string = "John"
```

### Scope rules

- `Local variables` are scoped to the function in which they are declared. They are not visible outside the function.
- `Package level variables` are scoped to the package in which they are declared. They are visible to all the functions in the package.
- `Exported variables` are scoped to the package in which they are declared. They are visible to all the functions in the package and other packages that import the package.

### Maps

- A map is an unordered collection of key-value pairs. Maps are similar to dictionaries in Python. The limitation of maps is that the key should be of the same type and the value should be of the same type. The key and value can be of any type.

### Structs

- A struct is a collection of fields. It is a data structure that lets us bundle together related data and behavior. We can use structs to represent real-world objects. It can handle multiple data types. It is similar to classes in other languages.

```go
type Person struct {
    name string
    age int
}
```

### Go routines

- A goroutine is a lightweight thread managed by the Go runtime. We can create a goroutine using the keyword `go`. It is similar to threads in other languages.

The purpose of a goroutine is to run a function concurrently with other functions. It is a function that is capable of running concurrently with other functions. It will not wait for the function to complete. It will execute the function concurrently.

```go
go func() {
    fmt.Println("Hello")
}()
```

- If the main function exits, the program will exit immediately even if the goroutine is still running. To prevent this, we can use the `WaitGroup` type. For Eg:

```go

#### WaitGroup

var wg sync.WaitGroup

func main() {
    wg.Add(1) // We are adding 1 to the WaitGroup
    go sayHello()
    wg.Wait() // We are waiting for the WaitGroup to become zero

}
func sayHello() {
    fmt.Println("Hello")
    wg.Done() // We are decrementing the WaitGroup by 1
}
```

`Add()` increments the WaitGroup counter by 1 and `Done()` decrements the WaitGroup counter by 1.