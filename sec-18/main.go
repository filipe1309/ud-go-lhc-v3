package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	flavors []string
}

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 18 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class128()
}

func class128() {
	fmt.Println("\nClass 128 - Hands-on exercise #53 - struct with slice")

	p1 := person{
		first: "John",
		last: "Doe",
		flavors: []string{"vanilla", "cream"},
	}

	p2 := person{
		first: "Elvis",
		last: "Presley",
		flavors: []string{"chocolate", "strwberry", "lemon"},
	}

	fmt.Println(p1)

	for _, v := range p1.flavors {
		fmt.Println(v)
	}

	fmt.Println(p2)

	for _, v := range p2.flavors {
		fmt.Println(v)
	}
}
