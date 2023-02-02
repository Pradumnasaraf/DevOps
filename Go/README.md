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
- [Let's go with golang](https://youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa) - Recommended
- [Golang Tutorial TechWorld with Nana](https://youtu.be/yyUHQIec83I)
- [Codeacademy Free course](https://www.codecademy.com/learn/learn-go)

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

### Mod

We can create a module using the following command. It will create a `go.mod` file. The use of `go.mod` is to track dependencies.

```bash
$ go mod init github.com/username/repo
```

### Build

We can build a program using the following command:

```bash
$ go build hello.go
```

we can also build for different platforms using the following command:

```bash
$ GOOS=linux GOARCH=amd64 go build hello.go
```

### Memory Allocation

- new() - It allocates memory. The size of the memory is equal to the size of the type. It returns a pointer to the memory.
- make() - It creates slices, maps, and channels only. It returns an initialized (not zeroed) value of type. It is used to create dynamically sized objects.

Garbage collection is a form of automatic memory management. The garbage collector frees the memory occupied by objects that are no longer in use by the program.

### Print

There are serveral ways to print in Go. We can use `fmt` package to print.

- `fmt.Println` - prints a line
- `fmt.Print` - prints without a new line
- `fmt.Printf` - prints with formatting

#### Placeholders for `printf` :

- %v - value in default format
- %+v - value in default format with field names
- %T - type of value
- %t - boolean
- %d - decimal integer
- %b - binary integer
- %c - character
- %x - hexadecimal integer
- %f - floating point number
- %s - string

Escape sequences:

- \n - newline
- \t - tab

### Variables

- `var` - variables with initializers.
- `const` - declares a constant value. Can't be updated.
- `:=` - short variable declaration. Can be used only inside a function. Called Syntax sugar. We can't use it to declare a global variable.

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

- When we declare a variable without a value, Go will assign a default value to it. Eg: `0` for `int`, `false` for `bool` etc.

- If we don't put a variable type go will automatically assign the type based on the value.-

### Scan

`fmt.Scan` reads text from standard input, scanning the text read into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

```go
fmt.Scan(&name)
```

#### Scanning though bufio

`bufio` package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.

```go
func main() {

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n') // It is text, err syntax. We put _ to ignore the error.
fmt.Println(text)
}
```

#### "text , err syntax"

In the above example we used `text, err` syntax. It is a common way to handle errors in Go. If we don't want to handle the error we can use `_` to ignore it.

```go
text, err := reader.ReadString('\n')

if err != nil {
    panic(err)
}
```

- `panic` - It



### Conversion

We can convert a value from one type to another. The expression T(v) converts the value v to the type T. We can use `strconv` package to convert a data type to another.


### Time

We can use `time` package to get the current time.

```go
t := time.Now()
fmt.Println(t)
```

### Pointers

A pointer holds the memory address of a value.

we can declare a pointer as follows:

```go
var p *int
```

We can get the memory address of a variable using `&` operator.

```go
var name = "John"
fmt.Println(&name) // it will print the memory address of the variable name
```

### Arrays

An array is a numbered sequence of elements of a single type with a fixed length. We can store a fixed size collection of elements of the same type.

We declare an array as follows:

```go
var arr [5]int  // array of 5 integers

arr[0] = 1
arr[1] = 2
```

### Slice

In this we don't need to specify the size of the array. It is a dynamically sized, flexible view into the elements of an array.

```go
var names []int  // slice of integers
```

Unlink arrays, we don't add elements to a slice using `arr[index] = value`. We use `append` function to add elements to a slice.

```go
names = append(names, 1)
names = append(names, 2)
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
names := []string{"John", "Paul", "George", "Ringo"}

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

- In Go, we don't need to write `break` after each case. It will automatically break after each case. If we want to execute the next case we can use `fallthrough` keyword.

```go
switch num {
case 1:
    fmt.Println("One")
    fallthrough // it will execute the next case
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

func main() {
    fmt.Println(add(1, 2)) // We can call the function by passing the arguments
}
```

- A function can return multiple values.

```go
func swap(x, y string) (string, string) { // When we have more than one return value we put them in parenthesis ().
return y, x
}

a, b := swap("hello", "world") // We can get the return values using multiple assignment
```

#### Anonymous functions

- We can declare a function without a name. Such functions are called anonymous functions.

```go
func(x, y int) int {
    return x + y
}
```

### Methods

- A method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name.

```go
type Person struct {
    name string
}

func (p Person) getName() string { // We can use the receiver argument to access the fields of the struct
    return p.name
}

func main() {
    p := Person{name: "John"}
    fmt.Println(p.getName())
}
```

### Defer

- A defer statement defers the execution of a function until the surrounding function returns.

```go
func main() {
    defer fmt.Println("world") // It will print "world" after the main function returns

    fmt.Println("hello")
}
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

```go
var cars = make(map[string]string, 0) // map of string to string

cars["Toyota"] = "Camry" // adding a key-value pair

delete(cars, "Toyota") // deleting a key-value pair
```

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


### goto

- The goto statement transfers control to the labeled statement. It is similar to the break statement in other languages.


### Math

- rand.Seed() - It is used to initialize the default Source to a deterministic state. If Seed is not called, the generator behaves as if seeded by Seed(1). It should only be called once. It is usually called before the first call to Intn or Float64.

```go
rand.Seed(time.Now().UnixNano())
```

### Json

- We can use the `json` package to encode and decode JSON data.

- `json.Marshal()` - It is used to encode a value to JSON. It returns a byte slice and an error.
- `json.MarshalIndent()` - It is used to encode a value to JSON with indentation. It returns a byte slice and an error.
- `json.Unmarshal()` - It is used to decode a JSON-encoded value. It returns an error.

In the Structs we can use the `json` tag to specify the name of the field in the JSON. I

```go

type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}    
``` 
