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

func main() {
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
}
