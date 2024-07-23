package main

import (
	"fmt"
	"strings"

	"github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/acdc"
	"github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/mystr"
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
	class251()
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

const s = `We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be?
Your playing small does not serve the world.
There is nothing enlightened about shrinking so that other people won't feel insecure around you.
We are all meant to shine, as children do.
We were born to make manifest the glory that is within us.
It's not just in some of us; it's in everyone.
And as we let our own light shine, we unconsciously give other people permission to do the same.
As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson`

func class251() {
	fmt.Println("\nClass 251 - Benchmarks examples")

	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\nCat: %v\n", mystr.Cat(xs))
	fmt.Printf("\nJoin: %v\n", mystr.Join(xs))
}
