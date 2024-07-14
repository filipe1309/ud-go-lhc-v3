package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 20 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class152()
	class153()
	class154()
	class155()
}

func class152() {
	fmt.Println("\nClass 152 - Hands-on exercise #57 - func concepts")
	x := sumNamedReturn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
	x = sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sumNamedReturn(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func class153() {
	fmt.Println("\nClass 153 - Hands-on exercise #58 - basic funcs")
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 21, "James Bond"
}

func class154() {
	fmt.Println("\nClass 154 - Hands-on exercise #59 - variadic func")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo2(xi...))
	fmt.Println(bar2(xi))
}

func foo2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func class155() {
	fmt.Println("\nClass 155 - Hands-on exercise #60 - defer func")
	fmt.Println("Before defer funcs")
	// LIFO
	defer fmt.Println("Deferred func 1")
	defer fmt.Println("Deferred func 2")
	defer fmt.Println("Deferred func 3")
	fmt.Println("After defer funcs")
}
