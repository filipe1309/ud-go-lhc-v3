package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 23 - Generics")
}

func main() {
	fmt.Println("Main function")
	class182()
	class183()
	class184()
	class185()
}

func class182() {
	fmt.Println("\nClass 182 - Type constraint")
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.1, 2.2))
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.1, 2.2))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

type myNumbers interface {
	int | float64
}

func addT2[T myNumbers](a, b T) T {
	return a + b
}

func class183() {
	fmt.Println("\nClass 183 - Type constraint & type set interface example")
	fmt.Println(addT2(1, 2))
	fmt.Println(addT2(1.1, 2.2))
}

type myAlias int

type myNumbers2 interface {
	~int | float64 // ~int means that myNumbers2 can accept any underlying type of int, or | myAlias could be used
}

func addT3[T myNumbers2](a, b T) T {
	return a + b
}

func class184() {
	fmt.Println("\nClass 184 - Type alias & underlying type constraint")
	// https://go.dev/blog/deconstructing-type-parameters
	var a myAlias = 1
	fmt.Println(addT3(a, 2))
	fmt.Println(addT3(1.1, 2.2))
}

type myNumbers3 interface {
	constraints.Integer | constraints.Float // constraints.Integer means that myNumbers3 can accept any underlying type of int, or | constraints.Float could be used
}

func addT4[T myNumbers3](a, b T) T {
	return a + b
}

func class185() {
	fmt.Println("\nClass 185 - Package constraint")
	fmt.Println(addT4(1, 2))
}
