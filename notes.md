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
```

`go help` command is used to get help for the go command.
```sh
go help
go help mod
```
