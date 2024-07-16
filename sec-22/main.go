package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 22 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class178()
	class179()
}

func class178() {
	fmt.Println("\nClass 178 - Hands-on exercise #74 - create a pointer")
	x := 42
	fmt.Printf("x value: %v, x type: %T, x address: %v\n", x, x, &x)
}


var (
	a, b, c *string
	d       *int
)


func class179() {
	fmt.Println("\nClass 179 - Hands-on exercise #75 - dereference an address")
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n

	fmt.Printf("a value: %v, a type: %T, a content: %v\n", a, a, *a)
	fmt.Printf("b value: %v, b type: %T, b content: %v\n", b, b, *b)
	fmt.Printf("c value: %v, c type: %T, c content: %v\n", c, c, *c)
	fmt.Printf("d value: %v, d type: %T,\t d content: %v\n", d, d, *d)
}

