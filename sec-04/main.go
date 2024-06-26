package main

import "fmt"

func main() {
	// a := 42
	// fmt.Println(a)

	// b, c, d, _, f := 0, 1, 2, 3, "happiness"
	// fmt.Println(b, c, d, f)

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
