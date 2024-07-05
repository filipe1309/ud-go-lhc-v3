package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 19 - Functions in the Go Programming Language")
}

func main() {
	fmt.Println("Main function")
	class133()
	class134()
	class135()
}

func class133() {
	fmt.Println("\nClass 133 - Syntax of functions in Go")
	foo()
	bar("Quiky")
	s := aloha("John")
	fmt.Println(s)
	dn, da := dogYears("Mel", 2)
	fmt.Println(dn, da)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("I am from bar, and my name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("I am from aloha, and my name is ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " dog years age is:"), age
}

func class134() {
	fmt.Println("\nClass 134 - Variadic parameter")
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is:", x)
}

func sum(ii ...int) int {
	fmt.Printf("%v \t %T\n", ii, ii)
	n := 0
	for _, v := range ii {
		n += v
	}

	return n
}

func class135() {
	fmt.Println("\nClass 135 - Unfurling a slice")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The sum is:", x)
}
