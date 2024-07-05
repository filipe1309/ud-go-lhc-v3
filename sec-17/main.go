package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 17 - Grouping data values - structs")
}

func main() {
	fmt.Println("Main function")
	class124()
}

func class124() {
	fmt.Println("\nClass 124 - Struct Introduction")
	p1 := person{
		first: "James",
		last: "Bond",
		age: 42,
	}
	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("Type: %T \t Value: %#v\n", p1, p1)
	fmt.Printf("p1.first: %v, p1.last: %v, p1.age: %v\n",p1.first, p1.last, p1.age)
}
