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
}

func class178() {
	fmt.Println("\nClass 178 - Hands-on exercise #74 - create a pointer")
	x := 42
	fmt.Printf("x value: %v, x type: %T, x address: %v\n", x, x, &x)
}
