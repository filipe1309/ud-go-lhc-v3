package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person // Embedded struct
	ltk bool
}

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 17 - Grouping data values - structs")
}

func main() {
	fmt.Println("Main function")
	class124()
	class125()
	class126()
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

func class125() {
	fmt.Println("\nClass 125 - Embedded structs")
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 42,
		},
		ltk: true,
	}
	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("Type: %T \t Value: %#v\n", sa1, sa1)
	fmt.Printf("sa1.first: %v, sa1.last: %v, sa1.age: %v, sa1.ltk: %v\n",sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Printf("sa1.person.first: %v, sa1.person.last: %v, sa1.person.age: %v, sa1.ltk: %v\n",sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk)
}

func class126() {
	fmt.Println("\nClass 124 - Anonymous structs")
	p1 := struct { // anonymous struct
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 42,
	}
	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
		age: 27,
	}

	fmt.Printf("Type: %T \t Value: %#v\n", p1, p1)
	fmt.Printf("Type: %T \t\t\t\t\t Value: %#v\n", p2, p2)
}
