---
sidebar_position: 1
title: Golang Introduction
description: Learn about Golang and its features.
tags: ["Golang", "Programming Language", "Go"]
keywords: ["Golang", "Programming Language", "Go"]
slug: "/golang"
---

Golang (or Go) is statically typed, compiled programming language designed at Google. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.

In Golang, everything is a package. A package is a collection of source files in the same directory that are compiled together. A package can be imported by other packages. `main` is a special package that defines a standalone executable program, not a library.

## Installation

- [Download](https://golang.org/dl/) and install Go
  
- Garbage collected
- Multithreading
- concurrency

## Packages

Module is a collection of related Go packages that are versioned together as a single unit.

- fmt - formatted I/O with functions analogous to C's printf and scanf.
- os - provides a platform-independent interface to operating system functionality.
- strconv - implements conversions to and from string representations of basic data types.
- time - provides functionality for measuring and displaying time.
- math - provides basic constants and mathematical functions.
- net/http - provides HTTP client and server implementations.
- encoding/json - implements encoding and decoding of JSON.


## Hello World - Running a program

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

## Go Mod (go.mod)

Go modules are a dependency management system that makes dependency version information explicit and easier to manage. Go modules are the future of dependency management in Go.

`go mod init github.com/username/repo` - creates a new module, initializing the go.mod file that describes it.

- `go mod tidy` - command will add any missing modules necessary to build the current module's packages and dependencies. It will also remove any unused modules that don't provide any relevant packages. It will update the go.mod file and the go.sum file.

- `go mod verify` - command will verify dependencies have expected content.

- `go list -m all` - command will list all modules needed to build the current module, as well as indirect and test dependencies.

- `go get` - command will add dependencies to current module and install them.

- `go list -m -versions <module name>` - command will list all available versions of a module.

- `go mod edit -go 1.16` - command will update the go directive in the go.mod file to the specified version.

- `go mod vendor` - command will copy all dependencies into a vendor directory.

## Sum (go.sum)

It is a file that contains the expected cryptographic checksums of the content of specific module versions. It ensures that dependencies have not been modified.

## Go Path

Go path is an environment variable that specifies the location of your workspace. It is used to find the location of your Go code.

## Build

We can build a program using the following command:

```bash
$ go build hello.go
```

we can also build for different platforms using the following command:

```bash
$ GOOS=linux GOARCH=amd64 go build hello.go
```

## Go Default Environment Variables

| **Env Name**        | **Description**                                                                                               | **Example**                                                                  |
|---------------------|---------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| `GO111MODULE`       | Controls Go module support. If empty, Go uses module mode when a `go.mod` file is present, otherwise falls back to GOPATH mode. | `GO111MODULE=on` (forces Go to use modules)                                  |
| `GOARCH`            | The architecture for which Go is building the binary (e.g., `amd64`, `arm64`).                                 | `GOARCH=arm64` (for Apple Silicon Macs)                                      |
| `GOBIN`             | Directory where `go install` places binaries.                                                                 | `GOBIN=$GOPATH/bin` (default if empty)                                        |
| `GOCACHE`           | Directory where Go stores build cache to speed up subsequent builds.                                          | `GOCACHE=/path/to/cache`                                                     |
| `GOENV`             | Path to the Go environment file storing Go environment variables.                                             | `GOENV=/path/to/go/env`                                                      |
| `GOEXE`             | File extension for Go executables (empty on Unix).                                                            | `GOEXE=.exe` (on Windows systems)                                             |
| `GOEXPERIMENT`      | Enables experimental features in Go.                                                                          | `GOEXPERIMENT=generics` (for future Go features)                             |
| `GOFLAGS`           | Flags passed to every `go` command.                                                                            | `GOFLAGS="-mod=readonly"` (prevents modifying `go.mod`)                       |
| `GOHOSTARCH`        | Architecture of the machine running the Go toolchain (same as `GOARCH`).                                      | `GOHOSTARCH=arm64`                                                           |
| `GOHOSTOS`          | OS of the machine running the Go toolchain (same as `GOOS`).                                                  | `GOHOSTOS=darwin`                                                            |
| `GOINSECURE`        | Specifies which modules can be fetched without HTTPS or checksum validation.                                 | `GOINSECURE=example.com/private` (for private, insecure repositories)       |
| `GOMODCACHE`        | Directory where downloaded Go modules are cached.                                                             | `GOMODCACHE=$GOPATH/pkg/mod`                                                 |
| `GONOPROXY`         | Module paths that bypass `GOPROXY`.                                                                            | `GONOPROXY=example.com/*` (bypass proxy for certain modules)                |
| `GONOSUMDB`         | Module paths that bypass checksum validation.                                                                 | `GONOSUMDB=example.com/*` (bypass checksum validation for private modules)  |
| `GOOS`              | The target operating system for the Go build (e.g., `linux`, `darwin`, `windows`).                             | `GOOS=linux` (to build a Linux binary)                                      |
| `GOPATH`            | The root directory for Go workspaces, storing source code, dependencies, and binaries.                        | `GOPATH=/Users/username/go` (default location)                               |
| `GOPRIVATE`         | Specifies module paths as private, bypassing proxies and checksum validation.                                 | `GOPRIVATE=example.com/*` (for private repositories)                         |
| `GOPROXY`           | Proxy servers used for fetching Go modules.                                                                    | `GOPROXY=https://proxy.golang.org` (default proxy)                           |
| `GOROOT`            | The directory where Go is installed (contains Go's standard library and tools).                               | `GOROOT=/opt/homebrew/Cellar/go/1.23.4/libexec`                              |
| `GOSUMDB`           | The checksum database for Go modules to verify downloaded content.                                             | `GOSUMDB=sum.golang.org`                                                     |
| `GOTMPDIR`          | Temporary directory for Go build-related files.                                                                | `GOTMPDIR=/path/to/tmp`                                                      |
| `GOTOOLCHAIN`       | Specifies the Go toolchain used for compiling code.                                                            | `GOTOOLCHAIN=go1.23` (to use a specific Go version toolchain)                |
| `GOTOOLDIR`         | Directory where Go's compiler tools (like linker and compiler) are located.                                    | `GOTOOLDIR=/opt/homebrew/Cellar/go/1.23.4/libexec/pkg/tool/darwin_arm64`     |
| `GOVCS`             | Controls versioning for Go modules' version control systems.                                                  | `GOVCS=git` (to use Git for version control)                                 |
| `GOVERSION`         | The current version of Go being used.                                                                           | `GOVERSION=go1.23.4`                                                         |
| `GODEBUG`           | Debugging options for Go runtime (e.g., for garbage collection or trace logs).                                 | `GODEBUG="gctrace=1"` (to trace garbage collection)                         |
| `GOTELEMETRY`       | Controls telemetry data collection.                                                                            | `GOTELEMETRY=on` (telemetry enabled)                                         |
| `GOTELEMETRYDIR`    | Directory where telemetry data is stored.                                                                      | `GOTELEMETRYDIR=/path/to/telemetry`                                          |
| `GCCGO`             | The path to the GCC Go compiler for CGo support.                                                               | `GCCGO=gccgo`                                                               |
| `GOARM64`           | Specifies the ARM64 architecture version.                                                                      | `GOARM64=v8.0`                                                              |
| `AR`, `CC`, `CXX`   | Path to the archiver, C compiler, and C++ compiler used for CGo.                                               | `AR=ar`, `CC=cc`, `CXX=c++`                                                 |
| `CGO_ENABLED`       | Enables or disables CGo, which allows Go to call C code.                                                       | `CGO_ENABLED=1` (enables CGo)                                               |
| `CGO_*` flags       | Flags passed to the C compiler and linker for CGo, like `CGO_CFLAGS` for compiling flags.                    | `CGO_CFLAGS="-O2"` (optimizes CGo code compilation)                         |
| `PKG_CONFIG`        | Specifies the path to the `pkg-config` tool for C libraries.                                                   | `PKG_CONFIG=pkg-config`                                                     |
| `GOGCCFLAGS`        | Flags passed to the GCC compiler for CGo-based Go code compilation.                                           | `GOGCCFLAGS="-fPIC -arch arm64 -pthread"`                                    |

## Memory Allocation

- new() - It allocates memory. The size of the memory is equal to the size of the type. It returns a pointer to the memory.
- make() - It creates slices, maps, and channels only. It returns an initialized (not zeroed) value of type. It is used to create dynamically sized objects.

Garbage collection is a form of automatic memory management. The garbage collector frees the memory occupied by objects that are no longer in use by the program.

## Print

There are serveral ways to print in Go. We can use `fmt` package to print.

- `fmt.Println` - prints a line
- `fmt.Print` - prints without a new line
- `fmt.Printf` - prints with formatting

Placeholders for `printf`

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

## Variables and Constants

- `var` - variables with initializers.
- `const` - declares a constant value. Can't be updated.
- `:=` - short variable declaration. Can be used only inside a function. Called Syntax sugar. We can't use it to declare a global variable.

## Data Types

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

### String

- A string is a sequence of characters. We can use double quotes to create a string.

```go
var name = "John"
```

- We can use backticks to create a raw string literal. It is used to create a string without escape sequences.

```go
var name = `John\n`
```

We can perform a lot of operations on strings like concatenation, length, indexing etc. Here are some examples:

```go
var name = "John"

fmt.Println(len(name)) // it will print the length of the string
fmt.Println(name[0]) // it will print the first character of the string
fmt.Println(name + " Doe") // it will concatenate the strings
```

With the `strings` package, we can perform more operations on strings like splitting, joining, replacing etc. It's one of most used packages in Go.

```go
import "strings"

var name = "John Doe"

fmt.Println(strings.Split(name, " ")) // it will split the string based on the space
fmt.Println(strings.Join([]string{"John", "Doe"}, " ")) // it will join the strings with a space
fmt.Println(strings.Replace(name, "John", "Jane", 1)) // it will replace the first occurrence of John with Jane
fmt.Println(strings.Contains(name, "John")) // it will check if the string contains John
fmy.Println(strings.ToLower(name)) // it will convert the string to lowercase
fmt.Println(strings.ToUpper(name)) // it will convert the string to uppercase
```

## Scan

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

### "text , err syntax"

In the above example we used `text, err` syntax. It is a common way to handle errors in Go. If we don't want to handle the error we can use `_` to ignore it.

```go
text, err := reader.ReadString('\n')

if err != nil {
    panic(err)
}
```

- `panic` - It is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.


## Conversion

We can convert a value from one type to another. The expression T(v) converts the value v to the type T. We can use `strconv` package to convert a data type to another.

```go
var num = 42
var stringNum = "42"

var str = strconv.Itoa(num) // it will convert the integer to a string
var intNum, _ = strconv.Atoi(stringNum) // it will convert the string to an integer
```

## Time

We can use `time` package to get the current time.

```go
t := time.Now()
fmt.Println(t)
```

## Pointers

A pointer is a variable that stores the memory address of another variable. We can declare a pointer by using `*` operator. Eg:

```go
var p *int
```

We can get the memory address of a variable using `&` operator.

```go
var name = "John"
fmt.Println(&name) // it will print the memory address of the variable name
```

We can get the value of a pointer using `*` operator. Called dereferencing.

```go

var name = "John"
var myName = &name

fmt.Println(*myName) // it will print the value of the variable name
```

- `*int` - the type `*int` is a pointer to an `int`.


```go
var name = "John"
fmt.Println(&name) // it will print the memory address of the variable name
```

## Arrays

An array is a numbered sequence of elements of a single type with a fixed length. We can store a fixed size collection of elements of the same type.

We declare an array as follows:

```go
var arr [5]int  // array of 5 integers

arr[0] = 1
arr[1] = 2
```

## Slice

In this we don't need to specify the size of the array. It is a dynamically sized, flexible view into the elements of an array.

```go
var names []int  // slice of integers
```

Unlink arrays, we don't add elements to a slice using `arr[index] = value`. We use `append` function to add elements to a slice.

```go
names = append(names, 1)
names = append(names, 2)
```

## Loops

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

## If Else

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

### Break and Continue

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

## Switch

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

### Fallthrough

In Go, we don't need to write `break` after each case. It will automatically break after each case. If we want to execute the next case we can use `fallthrough` keyword.

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

## Functions

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

### Anonymous functions

- We can declare a function without a name. Such functions are called anonymous functions.

```go
func(x, y int) int {
    return x + y
}
```

## Methods

A method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name. Receiver can be of any type. Receiver is a special type of parameter that is passed to a method. It is similar to `this` in other languages. It is used to access the fields and methods associated with the Type like a Struct. There are two types of receivers:

1. **Value Receiver**:  It is used when we don't want to modify the original value. 

> `func (t Test) printName() { fmt.Println(t.Name) }`

2. **Pointer Receiver**: It is used when we want to modify the original value. 

> `func (t *Test) printName() { fmt.Println(t.Name) }`


Here is an example of a method:

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

- A defer statement defers the execution of a function until the surrounding function returns.

```go
func main() {
    defer fmt.Println("world") // It will print "world" after the main function returns

    fmt.Println("hello")
}
```

## Mutex 

- Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.

```go
var mutex sync.Mutex

mutex.Lock() // It will lock the mutex
mutex.Unlock() // It will unlock the mutex
```

- RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer. When a writer is active, no readers can be active. 

```go
var rwMutex sync.RWMutex

rwMutex.RLock() // It will lock the mutex for reading
rwMutex.RUnlock() // It will unlock the mutex for reading

rwMutex.Lock() // It will lock the mutex for writing
rwMutex.Unlock() // It will unlock the mutex for writing
```

## Maps

- A map is an unordered collection of key-value pairs. Maps are similar to dictionaries in Python. The limitation of maps is that the key should be of the same type and the value should be of the same type. The key and value can be of any type.

```go
var cars = make(map[string]string, 0) // map of string to string

cars["Toyota"] = "Camry" // adding a key-value pair

delete(cars, "Toyota") // deleting a key-value pair
```

## Structs

- A struct is a collection of fields. It is a data structure that lets us bundle together related data and behavior. We can use structs to represent real-world objects. It can handle multiple data types. It is similar to classes in other languages.

```go
type Person struct {
    name string
    age int
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

#### WaitGroup

- A WaitGroup waits for a collection of goroutines to finish. The main function will wait for the goroutines to finish before exiting.

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


### Concurrency vs Parallelism

- **Concurrency** - It is the ability of a program to be decomposed into parts that can run independently of each other. It is the composition of independently executing processes. It is about dealing with lots of things at once.

- **Parallelism** - It is the ability of a program to run multiple tasks simultaneously. It is about doing lots of things at once.


## Math

- rand.Seed() - It is used to initialize the default Source to a deterministic state. If Seed is not called, the generator behaves as if seeded by Seed(1). It should only be called once. It is usually called before the first call to Intn or Float64.

```go
rand.Seed(time.Now().UnixNano())
```

## Json

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

## Channels

- A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. It is communication between goroutines. It is similar to pipes in other languages.

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

### Buffered Channel

A buffered channel is a channel with a buffer. It can hold multiple values. We can specify the buffer size when we create the channel.

```go
var ch = make(chan int, 5) // buffered channel with a buffer size of 5
```

### Unbuffered Channel

An unbuffered channel is a channel without a buffer. It can hold only one value. We can send a value to the channel only if there is a goroutine ready to receive the value.

```go
var ch = make(chan int) // unbuffered channel
```

### Closing a Channel

We can close a channel using the `close()` function. It is used to indicate that no more values will be sent on the channel. 

```go
msg := make(chan int)

go func() {
    msg <- 1
    close(msg) // After closing the channel we can't send any more values
}()
```

But here a catch even tho channel is closed we can still receive the values from it like zero, so it's dalmatic whether the it's channel is closed or the value is zero. To overcome this we can receive the value and a boolean value which will tell us whether the channel is closed or not.

```go
package main

import (
	"fmt"
	"sync"
)

var wait = sync.WaitGroup{}

func main() {

	myChannel := make(chan int, 1)

	wait.Add(2)
	go func() {
		fmt.Println("First Go Routine")
		wait.Done()

		hello, isChannelOpen := <- myChannel
		
		if !isChannelOpen {
			fmt.Println("Channel Closed")
		}
		fmt.Println(hello)
	}()

	go func() {
		fmt.Println("Second Go Routine")
		close(myChannel)
		wait.Done()
		// myChannel <- 5
	}()

	wait.Wait()
}
```

### Send Only Channel

```go
var ch = make(chan<- int) // send only channel
```

```go
go func (ch chan<- int) {
    ch <- 10
}(ch)
```

### Receive Only Channel

```go
var ch = make(<-chan int) // receive only channel
```

```go
go func (ch <-chan int) {
    val := <-ch
    fmt.Println(val)
}(ch)
```

## IIF's (Immediately Invoked Functions)

- An immediately invoked function is a function that is executed as soon as it is created. It is a function that is executed immediately after it is created. It is also known as a self-invoking function.

```go
func main() {
    func() {
        fmt.Println("Hello")
    }()
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
## goto

- The goto statement transfers control to the labeled statement. It is similar to the break statement in other languages. It is used to transfer control to a different part of the program.

```go
func main() {
    i := 0
    start:
    fmt.Println(i)
    i++
    if i < 5 {
        goto start
    }
}
```

## Scope rules

- `Local variables` are scoped to the function in which they are declared. They are not visible outside the function.
- `Package level variables` are scoped to the package in which they are declared. They are visible to all the functions in the package.
- `Exported variables` are scoped to the package in which they are declared. They are visible to all the functions in the package and other packages that import the package.

## Package level variables

- We can declare variables at the package level. They are called package level variables.

NOTE: We can't use the short variable declaration operator `:=` to declare package level variables.

```go
var name string = "John" // It will be available to all the functions in the package

func main() {
    fmt.Println(name)
}
```

## Expoting and Importing

We can export a function/ variable by capitalizing the first letter of the function/ variable name. Now we can use it in other packages.

```go
func Add(x, y int) int {
    return x + y
}

var Name string = "John"
```

## Code Organization

We can organize our code by putting functions/ variables in different files and can use them in the main file or calling the main function from other files.

Also, we can have multiple packages in a single directory.



## What's next?

- [Learning Resources](./learning-resources.md) - Learn more about Golang with these resources.
- [Other Resources](./other-resources.md) - A list of resources to learn more about Golang.
