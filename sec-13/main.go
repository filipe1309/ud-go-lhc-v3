package main

import (
	"fmt"
)

func main() {
	fmt.Println("Section 13 - Main function")
	class95()
}

func class95() {
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Printf("%v of type %T\n", a, a)

	b := [...]string{"James", "Bond", "Shaken, not stirred"}
	fmt.Printf("%v of type %T\n", b, b)
	// b = a // cannot use a (type [3]int) as type [3]string in assignment

	var c [2]int
	c[0] = 42
	c[1] = 43
	fmt.Printf("%v of type %T\n", c, c)

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
}
