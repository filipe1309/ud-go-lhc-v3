package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 22 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class178()
	class179()
	class180()
	class181()
}

func class178() {
	fmt.Println("\nClass 178 - Hands-on exercise #74 - create a pointer")
	x := 42
	fmt.Printf("x value: %v, x type: %T, x address: %v\n", x, x, &x)
}

var (
	a, b, c *string
	d       *int
)

func class179() {
	fmt.Println("\nClass 179 - Hands-on exercise #75 - dereference an address")
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n

	fmt.Printf("a value: %v, a type: %T, a content: %v\n", a, a, *a)
	fmt.Printf("b value: %v, b type: %T, b content: %v\n", b, b, *b)
	fmt.Printf("c value: %v, c type: %T, c content: %v\n", c, c, *c)
	fmt.Printf("d value: %v, d type: %T,\t d content: %v\n", d, d, *d)
}

type dog struct {
	first string
}

func (d dog) run() {
	fmt.Printf("%v is running\n", d.first)
}

func (d dog) walk() {
	fmt.Printf("%v is walking\n", d.first)
}

type yougin interface {
	run()
	walk()
}

func youngMove(y yougin) {
	y.walk()
	y.run()
}

func class180() {
	fmt.Println("\nClass 180 - Hands-on exercise #76 - interface implementation & method sets")
	d1 := dog{"Bob"}
	youngMove(d1)

	d2 := dog{"Alice"}
	youngMove(d2)
	d2 = d2.changeName("Alice2")
	youngMove(d2)
}

func (d dog) changeName(newName string) dog {
	d.first = newName
	return d
}

func class181() {
	fmt.Println("\nClass 181 - Hands-on exercise #77 - value & pointer semantics")
	p1 := person{"Alice"}
	fmt.Printf("p1: %v\n", p1)
	p1 = changeNameValue(p1, "Alice2")
	fmt.Printf("p1: %v\n", p1)

	p2 := person{"Bob"}
	fmt.Printf("p2: %v\n", p2)
	changeNamePointer(&p2, "Bob2")
	fmt.Printf("p2: %v\n", p2)
}

type person struct {
	first string
}

func changeNameValue(p person, n string) person {
	p.first = n
	return p
}

func changeNamePointer(p *person, n string) {
	p.first = n
}
