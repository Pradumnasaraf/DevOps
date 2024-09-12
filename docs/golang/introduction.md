---
sidebar_position: 1
title: Golang Introduction
---

Golang (or Go) is statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. In Golang, everything is a package. A package is a collection of source files in the same directory that are compiled together. A package can be imported by other packages. `main` is a special package that defines a standalone executable program, not a library.

## Installation

- [Download](https://golang.org/dl/) and install Go

## Modules and Packages

Modules are a collection of related Go packages that are versioned together as a single unit. For example `net/http` is a package that is part of the `net` module. Some of the standard library packages in Go are: 

- **fmt** - formatted I/O with functions analogous to C's printf and scanf.
- **os** - provides a platform-independent interface to operating system functionality.
- **strconv** - implements conversions to and from string representations of basic data types.
- **time** - provides functionality for measuring and displaying time.
- **math** - provides basic constants and mathematical functions.
- **net/http** - provides HTTP client and server implementations.
- **encoding/json** - implements encoding and decoding of JSON.

## Writing a program (skeleton)

> `main.go` file

```go
package main // package declaration

import "fmt" // import fmt package

import ( "fmt" "os") // import multiple packages

func main() {
fmt.Println("Hello, World!") // Println is a function from the fmt package
}
```
We can run the program using the following command:

```bash
$ go run main.go
```

## Go Mod 

Go Mod helps to manage dependencies in Go. It is a dependency management tool that is built into the Go toolchain. It is used to manage the dependencies of a Go project. It creates a `go.mod` file that describes the module's dependencies.

`go mod init github.com/username/repo` - creates a new module, initializing the go.mod file that describes it. In this case our desired module path is `github.com/username/repo`.

- `go mod tidy` - command will add any missing modules necessary to build the current module's packages and dependencies. It will also remove any unused modules that don't provide any relevant packages. It will update the go.mod file and the go.sum file.

- `go mod verify` - command will verify dependencies have expected content.

- `go list -m all` - command will list all modules needed to build the current module, as well as indirect and test dependencies.

- `go get` - command will add dependencies to current module and install them.

- `go list -m -versions <module name>` - command will list all available versions of a module.

- `go mod edit -go 1.16` - command will update the go directive in the go.mod file to the specified version.

- `go mod vendor` - command will copy all dependencies into a vendor directory.

## Sum (go.sum)

It is a file that contains the expected cryptographic checksums of the content of specific module versions. It ensures that dependencies have not been modified. It will automatically be created when we run `go mod tidy` or `go get`.  We don't need to worry about this file. It is automatically managed by Go.

## Build

We can build a Go program and create an executable file so that we can run from anywhere. We can use the `go build` command to build a Go program. It will create an executable file with the same name as the package name.

```bash
$ go build main.go 
$ go build . # it will build all the files in the current directory
$ go build -o app main.go # we can specify the name of the executable file
```

We can also build for different platforms using the following command:

```bash
$ GOOS=linux GOARCH=amd64 go build hello.go
```

## Memory Allocation

- new() - It allocates memory. The size of the memory is equal to the size of the type. It returns a pointer to the memory.
- make() - It creates slices, maps, and channels only. It returns an initialized (not zeroed) value of type. It is used to create dynamically sized objects.

It will be get clearer as we move forward.


### Print

There are several ways to print in Go. We generally use the `fmt` package to print in Go.

- `fmt.Println` - Prints a line with a new line
- `fmt.Print` - Prints without a new line
- `fmt.Printf` - Prints with formatting

Eg:

```go
func main() {
    fmt.Println("Hello, World!")
    fmt.Print("Hello, World!")
    fmt.Printf("Hello, %s!", "World")
}
```

#### Placeholders for `printf` :

- **%v** - value in default format
- **%+v** - value in default format with field names
- **%T** - type of value
- **%t** - boolean
- **%d** - decimal integer
- **%b** - binary integer
- **%c** - character
- **%x** - hexadecimal integer
- **%f** - floating point number
- **%s** - string

Eg:

```go
fmt.Printf("Hello, %s! You are %d years old.", "John", 25)
```

Escape sequences:

- \n - newline
- \t - tab
- \r - carriage return

Eg:

```go
func main() {
    fmt.Println("Hello, World!") // It will print Hello, World!
    fmt.Println("Hello, \tWorld!") // It will print Hello, World! with a tab
}
```

### Variables and Constants

- `var` - variables with initializers.
- `const` - declares a constant value. Can't be updated.
- `:=` - short variable declaration. Can be used only inside a function. Called Syntax sugar. We can't use it to declare a global variable.

```go
var name string = "John" // variable declaration with type
const pi = 3.14 // constant declaration

func main() {
    age := 25 // short variable declaration
}
```

### Scope rules

- `Local variables` are scoped to the function in which they are declared. They are not visible outside the function.
- `Package level variables` are scoped to the package in which they are declared. They are visible to all the functions in the package.
- `Exported variables` are scoped to the package in which they are declared. They are visible to all the functions in the package and other packages that import the package.
- `Block level variables` are scoped to the block in which they are declared. They are not visible outside the block.

NOTE: For  Package level variables We can't use the short variable declaration operator `:=` to declare package level variables.

```go
var name string = "John" // It will be available to all the functions in the package

func main() {
    fmt.Println(name)
}
```

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

## Scanner

### Scanning through fmt

`fmt.Scan` reads text from standard input, scanning the text read into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

```go
fmt.Scan(&name)
```

### Scanning though bufio

`bufio` package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.

```go
func main() {

reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n') // It is text, err syntax. We put _ to ignore the error.
fmt.Println(text)
}
```

## "text , err syntax" | ok, err syntax

In the above example we used `text, err` syntax. It is a common way to handle errors in Go. If we don't want to handle the error we can use `_` to ignore it.

```go
text, err := reader.ReadString('\n')

if err != nil {
    panic(err)
}
```

- `panic` - It is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.

## Conversion

We can convert a value from one type to another using the type conversion syntax. We can convert a value from one type to another using the type conversion syntax. Eg:

```go
var num int = 42
var numFloat float64 = float64(num)

var numFloat = 3.14
var numInt = int(numFloat)
```

Else we can also use the `strconv` package to convert a value to a string and vice versa.

```go
var num = 42
var numStr = strconv.Itoa(num)

var numStr = "42"
var num, _ = strconv.Atoi(numStr)
```

### Time

We can use `time` package to get the current time.

```go
t := time.Now()
fmt.Println(t)
```

### Math

We can use the `math` package to perform mathematical operations in Go. We can do operations like addition, subtraction, multiplication, division, etc.

```go
rand.Seed(time.Now().UnixNano())
```

## Strings Operations

We can perform various operations on strings in Go. Some of them are:

- `len()` - It is used to get the length of a string.
- `strings.ToUpper()` - It is used to convert a string to uppercase.
- `strings.ToLower()` - It is used to convert a string to lowercase.
- `strings.TrimSpace()` - It is used to remove leading and trailing spaces from a string.
- `strings.Split()` - It is used to split a string into substrings based on a separator.
- `strings.Contains()` - It is used to check if a string contains a substring.
- `strings.Replace()` - It is used to replace a substring with another substring.
- `strings.Join()` - It is used to join a slice of strings into a single string.
- `strings.Index()` - It is used to find the index of a substring in a string.

```go
func main() {

	var name = "John"
	fmt.Println(len(name))               // It will print the length of the string
	fmt.Println(strings.ToUpper(name))   // It will print the string in uppercase
	fmt.Println(strings.ToLower(name))   // It will print the string in lowercase
	fmt.Println(strings.Title(name))     // It will print the string in title case
	fmt.Println(strings.Repeat(name, 3)) // It will repeat the string 3 times
	fmt.Println(strings.Replace(name, "J", "K", 1)) // It will replace the first occurrence of J with K
	fmt.Println(strings.Split(name, "")) // It will split the string into a slice of characters
}
```

## Pointers

A pointer is a variable that stores the memory address of another variable. It is used to store the address of a variable. We can use the `&` operator to get the address of a variable and the `*` operator to get the value stored at the address.

Use of pointers:

- To pass a large struct to a function without copying it.
- To update the value of a variable in a function.

```go
myInt := 42
ptr := &myInt // & is used to get the address of the variable
// var ptr *int = &myInt

println("myInt:", ptr) // ptr is the address of the variable
println("myInt:", *ptr) // * is used to get the value stored at the address
```

- `*int` - the type `*int` is a pointer to an `int`.

## Arrays

An array is a collection of elements of the same type. It is a fixed-size collection of elements.

We declare an array as follows:

```go
var arr [5]int  // array of 5 integers
var arr1 = [5]int{1, 2, 3, 4, 5} // array of 5 integers with values

arr[0] = 1
arr[1] = 2
```

## Slice

Slice are more flexible than arrays. Their size is not fixed. A slice is a lightweight data structure that gives access to a subsequence of the elements of an array.

```go
var names []int  // slice of integers
var names = []int{1, 2, 3, 4, 5} // slice of integers with values
var names = make([]int, 5) // slice of integers with length 5
var names = [5]string{}
```

Unlink arrays, we don't add elements to a slice using `arr[index] = value`. We use `append` function to add elements to a slice.

```go
names = append(names, 1)
names = append(names, 2)
```

## Maps

Maps are used to store key-value pairs. It is an unordered collection of key-value pairs. Maps are similar to dictionaries in Python.

```go
var names = make(map[string]string) // map of string to string
var names = map[string]string{}
```

We can add key-value pairs to a map using the following syntax:

```go
names["John"] = "Doe"
names["Jane"] = "Doe"

fmt.Println(names["John"])
```

We can delete a key-value pair from a map using the `delete` function.

```go
delete(names, "John")
```

We can loop through a map using a `for` loop.

```go
for key, value := range names {
    fmt.Println(key, value)
}
```

## Structs

A struct is a collection of fields. It is a data structure that lets us bundle together related data and behavior. We can use structs to represent real-world objects.

```go
type Person struct {
    name string
    age int
}

var p Person
p.name = "John"
p.age = 25
```

## If else

In Go, the `if` statement is used to execute a block of code if a condition is true. The `else` statement is used to execute a block of code if the same condition is false.

```go
if num > 0 {
    fmt.Println("Positive")
} else if num == 0 {
    fmt.Println("Equal to zero")
} else {
    fmt.Println("Negative")
}
```

There is shorthand syntax for `if` statements. We can declare a variable and check the condition in the same line.

```go
if num := 10; num > 0 {
    fmt.Println("Positive")
}
```

## Loops

In Go, we have three types of loops:

- `for` - The for loop is the only loop statement in Go. It is similar to the for loop in other languages.
- `range` - The range form of the for loop iterates over a slice or map.
- `_` - The blank identifier is a special identifier that is used to ignore values when multiple values are returned from a function.

### For loop

For loop is used to execute a block of code multiple times. It is similar to the for loop in other languages.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### Range

The range form of the for loop iterates over a slice or map.

```go
names := []string{"John", "Paul", "George", "Ringo"}

for i, name := range names {
    fmt.Println(i, name)
}
```

### Blank Identifier

Blank identifier is a special identifier that is used to ignore values when multiple values are returned from a function.

We can use the blank identifier to ignore the index.

```go
for _, name := range names {
    fmt.Println(name)
}
```

Infinite loop

We can create an infinite loop using the `for` loop.

```go
for {
    fmt.Println("Hello")
}
```

## Switch

The switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

```go
switch num {
case 1 3: // It will execute the code if num is 1 or 3
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

We can use the `fallthrough` keyword to execute the next case. It will execute the next case even if the condition is false.

```go
switch num {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

We can also use the `switch` statement without a condition. It is similar to an if-else chain.

```go
switch {
case num > 0:
    fmt.Println("Positive")
case num == 0:
    fmt.Println("Equal to zero")
case num < 0:
    fmt.Println("Negative")
}
```

## Break and Continue

- `break` - Break is used to terminate the loop immediately. It is used to exit the loop. 

```go
for i := 0; i < 5; i++ {
    if i == 3 { // It will exit the loop when i is 3
        break
    }
    fmt.Println(i)
}
```

- `continue` - Continue is used to skip the current iteration of the loop. It is used to skip the current iteration and continue with the next iteration.

```go
for i := 0; i < 5; i++ {
    if i == 3 { // It will skip the iteration when i is 3
        continue
    }
    fmt.Println(i)
}
```

### Functions

Functions is a block of code that performs a specific task. It is a reusable piece of code. We can define a function using the `func` keyword.

```go

func hello() {
    fmt.Println("Hello")
}

func main() {
    hello()
}
```

### Functions with arguments

Functions can take arguments. We can specify the type of the arguments. We can pass them when we call the function.

```go
func add(x int, y int) { // We can specify the type of the parameters
    fmt.Println(x + y)
}

func main() {
    add(1, 2)
}
```

### Functions with return values

We can return a value from a function. We can specify the return type.

```go
func number() int { // We can specify the return type
    return 42
}
```

### Function with parameters and return values

We can have functions with parameters and return values.

```go
func add(x, y int) int { // We can specify the type of the parameters and the return type
    return x + y
}
```

When we have multiple multiple parameters we put them in parentheses and arguments are separated by commas. Remember Arguments are actual data.

```go
func addSub(x, y int) (int, int) { // We can specify the type of the parameters and the return type
    return x + y, x - y
}
```

### Variadic functions

A variadic function is a function that can accept a variable number of arguments. We can use the `...` operator to specify a variadic parameter.

```go
func sum(nums ...int) int { // We can specify the type of the parameters
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

### Anonymous functions

An anonymous function is a function that doesn't have a name. It is a function literal. We can assign it to a variable.

```go
add := func(x, y int) int {
    return x + y
}

fmt.Println(add(1, 2))
```

### IIFE (Immediately Invoked Function Expression)

An IIFE is a function that is immediately invoked. It is a function that is defined and called at the same time.

```go
func main() {
    func() {
        fmt.Println("Hello")
    }()
}
```

## Methods

Methods are functions that are associated with a type. They are similar to functions but are defined with a receiver. The receiver is like a parameter. It is the first argument of the method.

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

## Defer

Defer is used to delay the execution of a function until the surrounding function returns. It is used to ensure that a function call is performed regardless of the outcome of the surrounding function. It runs at the end of the function.

```go
func main() {
    defer fmt.Println("world") // It will print "world" after the main function returns

    fmt.Println("hello")
}
```

We can use defer to close a file after we open it.

```go
func main() {
    file, err := os.Open("file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close() // It will close the file after the main function returns
}
```

## File Handling

We can use the `os` package to work with files in Go. We can use the `os.Open` function to open a file. It returns a file and an error. We can use the `os.Create` function to create a file. It returns a file and an error.

```go

func main() {
    file, err := os.Open("file.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Read the file
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(data))
}
```

## Web Requests

We can use the `net/http` package to make web requests in Go. We can use the `http.Get` function to make a GET request. It returns a response and an error.

```go
func main() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(data))
}
```

## URLs

We can use the `net/url` package to parse URLs in Go. We can use the `url.Parse` function to parse a URL. It returns a URL and an error. We can construct or deconstruct URLs.

```go
const url = "https://example.com/path?name=John&age=25"

u, err := url.Parse(url)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(u.Scheme) // It will print the scheme
fmt.Println(u.Host) // It will print the host
fmt.Println(u.Path) // It will print the path
fmt.Println(u.RawQuery) // It will print the query
```

## Web Servers

We can create a web server in Go using the `net/http` package. We can use the `http.HandleFunc` function to register a handler function for a pattern. We can use the `http.ListenAndServe` function to start the server.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

This is provided by the `net/http` package, but we can use frameworks like `Gin` or `Echo` to create web servers.

### JSON

We can use the `encoding/json` package to encode and decode JSON data in Go. We can use the `json.Marshal` function to encode a value to JSON. It returns a byte slice and an error. We can use the `json.Unmarshal` function to decode a JSON-encoded value. It returns an error.

```go
type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func main() {
    p := Person{Name: "John", Age: 25}

    data, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(data))

    var p1 Person
    err = json.Unmarshal(data, &p1)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(p1)
}
```


## Go routines and WaitGroup

- A goroutine is a lightweight thread managed by the Go runtime. We can create a goroutine using the keyword `go`. It is similar to threads in other languages. The purpose of a goroutine is to run a function concurrently with other functions. 

```go
go func() {
    fmt.Println("Hello")
}()
```

- If the main function exits, the program will exit immediately even if the goroutine is still running. To prevent this, we can use the `WaitGroup` type. For Eg:

```go

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


## Race Conditions and Mutex

Race conditions occur when two or more goroutines access the same variable concurrently and at least one of the accesses is a write. It can lead to unpredictable behavior. To overcome this we can use the `sync` package. We can use the `sync.Mutex` type to lock and unlock a variable.

```go
func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("One Routine")
		mut.Lock() // Lock the memory location of score
		score = append(score, 1)
		mut.Unlock() // Unlock the memory location of score
		wg.Done()
	}(wg, mut) // Pass the pointer of WaitGroup and Mutex
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
```

- `Lock()` - It will lock the mutex and prevent other goroutines from accessing the variable
- `Unlock()` - It will unlock the mutex and allow other goroutines to access the variable


- `RWMutex` - It is a reader/writer mutual exclusion lock. It can be locked for reading or writing. It allows multiple readers or a single writer at a time.

```go
var rwMutex sync.RWMutex

rwMutex.RLock() // It will lock the mutex for reading
rwMutex.RUnlock() // It will unlock the mutex for reading
```

## Channels

A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. It is communication between goroutines. It is similar to pipes in other languages.

- `make()` - It is used to create a channel. It takes the type of the channel as an argument.

```go
ch := make(chan int)

go func() {
    ch <- 10
}()

val := <-ch
fmt.Println(val)
```

We can create a buffered channel by passing the buffer size as the second argument to the `make()` function. By default, the channel is unbuffered and can only hold one value. So, if we try to send multiple value to the channel it will give an error.

```go
var ch = make(chan int) // unbuffered channel

var ch = make(chan int, 10) // buffered channel
```

```go
ch := make(chan int)

ch <- 10 // It will send 10 to the channel

<- ch // It will receive from the channel
val, ok := <- ch // It will receive from the channel and check if the channel is closed or not
```

- `close()` - It is used to close a channel. It takes the channel as an argument.


#### Send Only Channel

```go
var ch = make(chan<- int) // send only channel
```

Receive Only Channel

```go
var ch = make(<-chan int) // receive only channel
```

## Interfaces

An interface is a collection of method signatures. It is a set of methods that a type must implement. We can use interfaces to define the behavior of an object. We can use the `interface` keyword to define an interface.

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}
```

We can use interfaces to create polymorphic behavior. We can use an interface as a type.

```go
func getArea(s Shape) float64 {
    return s.Area()
}

func main() {
    c := Circle{radius: 5}
    fmt.Println(getArea(c))
}
```
## Error Handling

In Go, errors are values. We can use the `error` type to represent an error. We can use the `errors.New` function to create a new error. It returns an error.

```go
func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
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

### goto

- `goto` is a statement that allows us to jump to a labeled statement in the same function. It is similar to the `goto` statement in other languages.

```go
func main() {
    i := 0
loop:
    fmt.Println(i)
    i++
    if i < 5 {
        goto loop
    }
}
```

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about Golang with these resources.
- [Other Resources](./other-resources.md) - A list of resources to learn more about Golang.
- 