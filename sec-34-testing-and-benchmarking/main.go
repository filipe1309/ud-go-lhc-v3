package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 34 - Testing & benchmarking")
}

func main() {
	fmt.Println("Main function")
	class245()
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func class245() {
	fmt.Println("\nClass 245 - Introduction")

	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 9 + 1 =", mySum(5, 9, 1))
}
