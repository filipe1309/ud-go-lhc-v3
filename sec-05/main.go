package main

import (
	"fmt"
)

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	c3 = iota // c0 == 0
	c4
	c5
	c6
)

type ByteSize int

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1024 Bytes
	MB                             // 1024 KB
	GB                             // 1024 MB
	TB                             // 1024 GB
	PB                             // 1024 TB
	EB                             // 1024 PB
	// ZB // 1024 EB
	// YB // 1024 ZB
)

func main() {
	/*======================= Class 24 =======================*/
	fmt.Println("\nClass 24")

	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)

	const (
		_ = iota
		one
		two
		three
		four
		five
		six
	)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", one, one)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<one, 1<<one)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<two, 1<<two)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<three, 1<<three)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<four, 1<<four)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<five, 1<<five)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
	fmt.Printf("%d \t %b\n", 1<<six, 1<<six)

	/*======================= Class 25 =======================*/
	fmt.Println("\nClass 25")

	fmt.Printf("KB dec: %d \t\t\t bin: %b\n", KB, KB)
	fmt.Printf("MB dec: %d \t\t bin: %b\n", MB, MB)
	fmt.Printf("GB dec: %d \t\t bin: %b\n", GB, GB)
	fmt.Printf("TB dec: %d \t\t bin: %b\n", TB, TB)
	fmt.Printf("PB dec: %d \t bin: %b\n", PB, PB)
	fmt.Printf("EB dec: %d \t bin: %b\n", EB, EB)
	// fmt.Printf("ZB: %d \t %b\n", ZB, ZB)
	// fmt.Printf("YB: %d \t %b\n", YB, YB)

	/*
		PB		    1125899906842624
		EB		 1152921504606846976
		int		18446744073709551615
	*/

	/*======================= Class 26 =======================*/
	// zero value, :=, type specificity, blank identifier
	fmt.Println("\nClass 26")

	// var for zero value
	var zero int
	fmt.Println(zero)

	// := for short declaration
	short := 42
	fmt.Println(short)

	// multiple initializations
	var first, second int = 1, 2
	fmt.Println(first, second)

	// var when specificity is required
	var third float32 = 3.0
	fmt.Printf("%v, %T\n", third, third)

	// blank identifier
	_, fourth := 4, 5
	fmt.Println(fourth)

	/*======================= Class 27 =======================*/
	// Hands-on exercise #11 - printf verbs to show values and types
	fmt.Println("\nClass 27")

	myString, myInt, myFloat := "Hello, World!", 42, 42.42
	fmt.Printf("Value: %v \t Type: %T\n", myString, myString)
	fmt.Printf("Value: %v \t\t Type: %T\n", myInt, myInt)
	fmt.Printf("Value: %v \t\t Type: %T\n", myFloat, myFloat)

	/*======================= Class 28 =======================*/
	// Hands-on exercise #12 - Printf binary, decimal & hexadecimal
	fmt.Println("\nClass 28")

	valOne, valTwo, valThree := 747, 911, 90210
	fmt.Printf("Decimal: %d \t Binary: %b \t\t Hexadecimal: %#X\n", valOne, valOne, valOne)
	fmt.Printf("Decimal: %d \t Binary: %b \t\t Hexadecimal: %#X\n", valTwo, valTwo, valTwo)
	fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal: %#X\n", valThree, valThree, valThree)

	/*======================= Class 29 =======================*/
	// Hands-on exercise #13 - Signed and Unsigned integers
	fmt.Println("\nClass 29")
	var mySignedVar int8 = 127
	fmt.Printf("mySignedVar \t Decimal: %d \t Binary: %b\n", mySignedVar, mySignedVar)
	var myUnsignedVar uint8 = 255
	fmt.Printf("myUnsignedVar \t Decimal: %d \t Binary: %b\n", myUnsignedVar, myUnsignedVar)

}
