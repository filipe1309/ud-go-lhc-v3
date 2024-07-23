package main

import (
	"fmt"

	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/dog"
	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/mymath"
	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/quote"
	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/word"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 35 - Exercises - Ninja Level 13")
}

func main() {
	fmt.Println("Main function")
	class253()
	class254()
	class255()
}

type canine struct {
	name string
	age  int
}

func class253() {
	fmt.Println("\nClass 253 - Hands-on exercise #1")
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(10))
}

func class254() {
	fmt.Println("\nClass 254 - Hands-on exercise #2")

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}

func class255() {
	fmt.Println("\nClass 255 - Hands-on exercise #3")

	xxi := gen()
	for _, xi := range xxi {
		fmt.Println(mymath.CenteredAvg(xi))
	}
}
