package main

import "fmt"

func main() {
	// Class 13

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

	// Class 14

	// adams := 42
	// fmt.Printf("42 as binary, %b \n", adams)
	// fmt.Printf("42 as hexadecimal, %x \n", adams)

	// // print these values as both binary & hexadecimal
	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", a, a, a)
	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", b, b, b)
	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", c, c, c)
	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", d, d, d)
	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", e, e, e)
	// fmt.Printf("dec: %d \t bin: %b \t hexa: %#X\n", f, f, f)

	// Class 16
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
