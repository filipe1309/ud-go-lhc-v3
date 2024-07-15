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
}

func class169() {
	fmt.Println("\nClass 169 - Waht are pointers?")
	x := 42
	fmt.Println(x) // x value
	fmt.Println(&x) // x address
}
