package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 21 - Pointers")
}

func main() {
	fmt.Println("Main function")
	class169()
	class170()
}

func class169() {
	fmt.Println("\nClass 169 - What are pointers?")
	x := 42
	fmt.Println(x) // x value
	fmt.Println(&x) // x address
}

func class170() {
	fmt.Println("\nClass 170 - Seeing type & value for pointers?")
	x := 42
	fmt.Printf("x value: %v, x type: %T\n", &x, &x)

	y := "James Bond"
	fmt.Printf("y value: %v, y type: %T\n", &y, &y)
}
