# Notes

> notes taken during the course

## Section 1

https://go.dev/doc/effective_go

https://github.com/GoesToEleven/learn-to-code-go-version-03

https://forum.golangbridge.org/

https://go.dev/play/

## Section 2

### Hello World

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers 🚀🧙‍♂️")
}
```

https://go.dev/play/p/joUNGupswT9


Standard library
https://pkg.go.dev/std
https://pkg.go.dev/fmt@go1.22.4

https://go.dev/play/p/4tquc1jpZXQ

```go
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	//fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%v is %v years old. \t And the type is %T and %T", name, age, name, age)
}
```
https://go.dev/play/p/fAQ1Q7SRHXX


https://go-proverbs.github.io/


### ASCII, Unicode & UTF-8
https://go.dev/play/p/PnzhfYhGWCC

### Raw string literals

```go
package main

import "fmt"

func main() {

	fmt.Println("Hello Gophers! ❤️💕😊👍😁(❁´◡`❁)£¥©🙌👌🎶😎🐼🦄🦁🐶😺🤓")

	// raw string literal
	fmt.Println(`
	UTF-8 saves space. 
	In UTF-8, common characters like “C” take 8 bits, 
	while rare characters like “💕” take 32 bits. 
	Other characters take 16 or 24 bits. 
	A blog post like this one takes about 
	four times less space in UTF-8 
	than it would in UTF-32. 
	So it loads four times faster.
	`)

	// https://go.dev/ref/spec#String_literals
}
```

https://go.dev/play/p/wge8tWLEQNM

## Section 4


https://go.dev/play/p/C5k8w07Fh7Y

https://go.dev/play/p/uvggBjVF47O

https://en.wikipedia.org/wiki/Numeral_system

https://en.wikipedia.org/wiki/Primitive_data_type

https://pkg.go.dev/builtin

https://go.dev/blog/constants

We DECLARE a VARIABLE of a certain TYPE
- is can only hold VALUES of that TYPE
- Go is statically typed

## Section 5

https://go.dev/tour/list

https://go.dev/play/p/jZUmqIhqaIC

https://go.dev/play/p/RoPv2HjoALg


## Section 6

> "Persistently, Patiently you are bound to succeed" quote by Todd McLeod

## Section 7

`go mod init` command is used to initialize a new module, creating a new go.mod file in the current directory.
```sh
go mod init github.com/filipe1309/ud-go-lhc-v3
```

`go mod tidy` command is used to add missing and remove unused modules.
```sh
go mod tidy
```

`go run` command compiles and runs the program.
```sh
go run main.go
# OR
go run ./... # run all files in the current directory
```

`go build` command compiles the program into an executable file.
```sh
go build
```

Cross-compiling

```sh
go env GOARCH GOOS # to see the current architecture and OS, ex: arm64 darwin
GOOS=windows go build -o myprogram.exe
GOOS=linux go build
GOOS=darwin go build
```

Environment variables

```sh
go env # to see all environment variables
go env GOROOT # to see the Go root directory
go env GOPATH # to see the Go path directory
```

`go install` command compiles and installs the package in the `bin` directory.
```sh
go install
```

## Section 8 - go mod and dependency management

`go get` command is used to download packages and dependencies.
```sh
go get github.com/filipe1309/ud-go-lhc-v3

go get github.com/filipe1309my-go-library@latest
```

`go help` command is used to get help for the go command.
```sh
go help
go help mod
```

"When thr bird and the book disagree, believe the bird" quote by James Audubon

## Section 9 - Hands-on exercises

`go install` command compiles and installs the package in the `$GOPATH/bin` directory.
```sh
go install
```

```sh
go get github.com/GoesToEleven/puppy@v1.2.0 # get a specific version
go get github.com/GoesToEleven/puppy@latest # get the latest version or update to the latest version
```
## Section 11 - Control flow

- Sequence
- Conditional
- Loop

The Stack = LIFO (Last In First Out), is faster and requires less overhead than the Heap.
The Heap = is slower and requires more overhead than the Stack, but is more flexible.

Go programs try to allocate as much memory as possible on the stack, because it is faster and requires less overhead.
However, whaen a program needs to store a large amount of data, or data with a longer lifetime than the stack, it must use the heap.

"comma ok idiom" is used to check if a value is present in a map.
```go
if value, ok := myMap[key]; ok {
  fmt.Println(value)
}
```

`select` statement is used to wait on multiple channels for concurrent communication.
```go
channel1 := make(chan string)
channel2 := make(chan string)

go func() {
  time.Sleep(1 * time.Second)
  channel1 <- "one"
}()

go func() {
  time.Sleep(2 * time.Second)
  channel2 <- "two"
}()

select {
  case msg1 := <-channel1:
    fmt.Println("received", msg1)
  case msg2 := <-channel2:
    fmt.Println("received", msg2)
  default:
    fmt.Println("no activity")
}
```

`for range` loop is used to iterate over a slice, array, string, map, or channel.
```go
// Slice example
mySlice := []int{1, 2, 3, 4, 5} // slice of integers
for index, value := range mySlice {
  fmt.Println(index, value)
}

// Map example
myMap := map[string]int{
  "a": 1, 
  "b": 2, 
  "c": 3
} // map of strings to integers
for key, value := range myMap {
  fmt.Println(key, value)
}
```

## Section 12 - Hands-on exercises

Checksum
```sh
shasum -a 256 sec-12/SNOWY-EVENING.txt
# 7c6c8937b2a120af15849db05c9f46326761e0eec852a2e973b1e0b6acd59a01  sec-12/SNOWY-EVENING.txt
```

## Section 13 - Grouping data vaues - array & slice

https://go.dev/play/p/f-7aufl2DO

https://go.dev/play/p/_tZD1mCUjej


Everything is passed by value in Go.

https://go.dev/play/p/ZBz2jFnL8oc


## Secttion 17 - Grounping data values - structs

### Composition

Composition is a way to combine multiple types into a single type, creating a new type that has all the fields and methods of the types it is composed of.

- It's all about types
 - Go has OOP aspects. But is all about TYPE.
- Go is Object Oriented
 - Encapsulation: state (fields) + behavior (methods) + exported/unexported
 - Reusability: inheritance (embedded types)
 - Polymorphism: interfaces
 - Overriding: promotion

 In Go:
 - We don't have `classes`, we have `types`
 - We don't have `inheritance`, we have `composition`
 - We don't have `instance of a class`, we have `value of a type`
 - We don't have `super`, we have `promotion`
 - We don't have `method overloading`, we have `method overriding`

 User defined types
 - We can declare our own types
 - We can declare types based on existing types
 - Why?
	- To make our code more descriptive
	- To make our code more type-safe
	- If we need to attach methods to a type


Go is about easy of programming.

Named types vs Anonymous types
- Anonymous types are indeterminate
 - They have not been declared as a type yet
 - The compiler will infer the type
 - Ex: `var x int`
- Named types are determinate
 - They have been declared as a type
 - The compiler will not infer the type
 - Ex: `type hotdog int`
- Padding & architectural alignment conventions
 - Logically organize your fields together
 - Readability & clarity trump performance as a design concern
 - Go will be performant. Go fo readability first. However, if you need to optimize: lay the fields out in the order of largest to smallest, eg: `int64` first, then `int32`, then `int16`, then `int8`

 ## Section 19 - Functions in the Go Programming Language

 Define with `parameters`
 Call with `arguments`

 Everything in Go is pass by `value`, which means that when we pass a value to a function, a copy of the value is passed.



