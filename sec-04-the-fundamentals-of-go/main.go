package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 4 - The fundamentals of Go")
}

func main() {
	class13()
	class14()
	class16()
}

func class13() {
	fmt.Println("\nClass 13 - Variables, zero values, blank identifier")

	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	/*
		var g int
		fmt.Println(g)
		g = 43
		fmt.Println(g)

		// declare a variable to hold a VALUE of a certain TYPE
		// initializing a variable
		var h int = 44
		fmt.Println(h)
	*/
}

func class14() {
	fmt.Println("\nClass 14 - Using printf for decimal & hexadecimal values")

	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", a, a, a)
	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", b, b, b)
	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", c, c, c)
	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", d, d, d)
	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", e, e, e)
	fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", f, f, f)
}

func class16() {
	fmt.Println("\nClass 16 - Values, types, conversion, scope & housekeeping")
	y := 42
	z := 42.0
	fmt.Printf("%v of type %T \n", y, y)
	fmt.Printf("%v of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	// this does work
	z = float64(m)
	fmt.Printf("%v of type %T \n", z, z)
}
