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
