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

 Everything in Go is pass by `value` (copy), which means that when we pass a value to a function, a copy of the value is passed.

```go
func (r receiver) identifier(parameters) (return(s)) {
	// code
}
```

### Interfaces & polymorphism

An interface in Go defines a set of methods that a type must implement in order to satisfy the interface.

Polymorphism is the ability of a VALUE of a certain TYPE to be treated as a VALUE of another TYPE.

Values can be of more than one type in Go.

## Section 20 - Hands-on exercises

### Documentation
Class 161. Hands-on exercise #66 - documenting code with comments

```sh
go doc -cmd -all # to see all the documentation
```

1. Package documentation
```sh
// Package dog converts human years to dog years.
package dog
```

2. Function documentation
```sh
// Years converts human years to dog years.
func Years(humanYears int) int {
	return humanYears * 7
}
```

3. Type documentation
```sh
// Dog is a type that represents a dog.
type Dog struct {
	Name string
	Age  int
}
```

4. Const and variable documentation
```sh
// DogYears is a constant that represents the number of dog years in a human year.
const DogYears = 7
```

### Interfaces & mocking testing a db

https://go.dev/play/p/mpDK30B1n5l

Pointers are used in method receivers in Go for two main reasons:

1. Efficiency: When a large struct is passed by value to a method (i.e., without using a pointer), the entire struct must be copied, which can be inefficient in terms of memory and performance. When a pointer to the struct is passed instead, only the memory address of the struct is passed, which is typically much more efficient.

2. Modification: If you want your method to modify the state of the receiver, the receiver must be a pointer. This is because Go is a pass-by-value language, meaning that when a value is passed to a function, a copy of that value is made in the function's scope. If the function modifies the copy, the original value is not affected. But if a pointer to the value is passed, the function can modify the original value.

In the case of your code, both `Service` and `MockDatastore` methods are defined with pointer receivers because the methods may need to modify the state of the receiver. For example, `MockDatastore.SaveUser` needs to add a user to the `Users` map of the `MockDatastore`. If `md` were not a pointer, `SaveUser` would receive a copy of the `MockDatastore`, and changes would only be made to that copy, not the original `MockDatastore`.

```go
func (md *MockDatastore) SaveUser(u User) error {
    _, ok := md.Users[u.ID]
    if ok {
        return fmt.Errorf("User %v already exists", u.ID)
    }
    md.Users[u.ID] = u
    return nil
}
```

Also, note that even when methods only read the receiver's data, you might still want to use a pointer receiver for efficiency reasons, especially if the struct is large.

A map is just a hash table.
The data is arranged into an array of buckets.
...
When the hashtable grows, we allocate a new array
of buckets twice as big. Buckets are incrementally
copied from the old bucket array to the new bucket array.

https://cs.opensource.google/go/go/+/refs/tags/go1.20.4:src/runtime/map.go

// A header for a Go map.
```go
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/reflectdata/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

## Section 21 - Pointers

### Pointers, values, the stack & the heap

- Everything in Go is pass by value
- When we pass a value to a function, a copy of the value is passed
- Pointers allow us to share a value stored in memory with a function

Value semantics & the stack: In Go, when we pass a value to a function, a copy of the value is passed. This is known as value semantics. The value is stored on the stack, which is a region of memory that is fast and has low overhead. The stack is used for local variables and function calls.

Pointers & the heap: When we pass a pointer to a function, the memory address of the value is passed. This is known as reference semantics. The value is stored on the heap, which is a region of memory that is slower and has higher overhead. The heap is used for dynamically allocated memory.

Escape analysis: The Go compiler performs escape analysis to determine whether a value can be allocated on the stack or must be allocated on the heap. If a value escapes to the heap, it is said to escape to the heap.

```sh
go run -gcflags "-m" main.go # to see the escape analysis
```

## Section 23 - Generics

https://go.dev/doc/tutorial/generics

Concrete types: A type that is fully defined and can be instantiated. 
Is a type you can directtly instatiate or create a value from. 
For example, `int`, `string`, and `float64` are concrete types.

```go
type Employee struct {
	Name string
	Age  int
}
emp := Employee{Name: "John Doe", Age:  30}
```

Interface types: A type that defines a contract, or set of methods, that a concrete type must implement in order to satisfy the interface.
Interface types are abstract, they represent behavior or type but not a specific set of values.

```go
type Stringer interface {
	String() string
}
```

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

## Section 24 - Application

https://mholt.github.io/json-to-go/

## Section 26 - Concurrency

**Do not communicate by sharing memory; instead, share memory by communicating.**

Concurrency is the ability to run multiple tasks at the same time.

Parallelism is the ability to run multiple tasks simultaneously.

The method set of a type determines the interfaces that the type implements.

Race conditions occur when two or more goroutines access shared data and attempt to modify it at the same time.

The `go run -race` command is used to verify that:
```sh
go run -race sec-26/main.go
```


## Section 28 - Channels

Channels block!!!

Channels are a way for goroutines to communicate with each other and synchronize their execution.

Range over a channel will block until the channel is closed.

## Section 32 - Writing documentation

- `go doc <pkg>`
- `go doc <sym>[.<method>]`
- `go doc [<pkg>.]<sym>[.<method>]`
- `go doc [<pkg>.][<sym>.]<method>`

```sh
go doc fmt # to see the documentation for the fmt package
```

```sh
go doc fmt.Println # to see the documentation for the fmt.Println function
```

https://pkg.go.dev/github.com/filipe1309/my-go-library

```sh
go doc github.com/filipe1309/my-go-library # to see the documentation for the my-go-library package
```


## Section 34 - Testing & Benchmarking

- `go test`
- `go test -v`
- `go test -cover`
- `go test -coverprofile c.out`
- `go tool cover -html=c.out`
- `go test -bench .`
- `go test -benchmem`
- `go test -bench . -benchtime 3s`

### Remember to BET

- Benchmark
- Example
- Test

```go
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}
```

```go
func ExampleHello() {
	fmt.Println(Hello())
	// Output: Hello, Gophers!
}
```

```go
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, Gophers!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

## Section 43 - Error handling 
> also in Section 30

"Error values in Go aren't special, they are just values like any other, and so you have the entire language at your disposal to deal with them." - Rob Pike

```go
if err != nil {
	log.Fatal(err)
}
```

```go
if err != nil {
	log.Fatalf("an error occurred: %v", err)
}
```

```go
if err != nil {
	log.Panic(err)
}
```

```go
if err != nil {
	log.Panicf("an error occurred: %v", err)
}
```

```go
if err != nil {
	log.Println(err)
}
```

```go
if err != nil {
	log.Printf("an error occurred: %v", err)
}
```

https://go.dev/doc/effective_go#errors



