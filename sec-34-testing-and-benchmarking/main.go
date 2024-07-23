package main

import (
	"fmt"

	"github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/acdc"
	"github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/saying"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 34 - Testing & benchmarking")
}

func main() {
	fmt.Println("Main function")
	class245()
	class247()
	class249()
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

func class247() {
	fmt.Println("\nClass 247 - Example tests")

	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}

func class249() {
	fmt.Println("\nClass 249 - Benchmark")

	fmt.Println(saying.Greet("James"))
}
