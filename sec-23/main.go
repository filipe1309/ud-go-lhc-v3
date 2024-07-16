package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 23 - Generics")
}

func main() {
	fmt.Println("Main function")
	class182()
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
