package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

type human interface {
	speak()
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am",  p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent",  sa.first)
}

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 19 - Functions in the Go Programming Language")
}

func main() {
	fmt.Println("Main function")
	class133()
	class134()
	class135()
	class136()
	class137()
	class138()
	class139()
	class140()
	class141()
}

func class133() {
	fmt.Println("\nClass 133 - Syntax of functions in Go")
	foo()
	bar("Quiky")
	s := aloha("John")
	fmt.Println(s)
	dn, da := dogYears("Mel", 2)
	fmt.Println(dn, da)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("I am from bar, and my name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("I am from aloha, and my name is ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " dog years age is:"), age
}

func class134() {
	fmt.Println("\nClass 134 - Variadic parameter")
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is:", x)
}

func sum(ii ...int) int {
	fmt.Printf("%v \t %T\n", ii, ii)
	n := 0
	for _, v := range ii {
		n += v
	}

	return n
}

func class135() {
	fmt.Println("\nClass 135 - Unfurling a slice")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The sum is:", x)
}

func class136() {
	fmt.Println("\nClass 136 - Defer")
	defer foo()
	bar("Quiky")
}

func class137() {
	fmt.Println("\nClass 137 - Methods")
	p1 := person{
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}
	p1.speak()
	p2.speak()
}

func class138() {
	fmt.Println("\nClass 138 - Interfaces & polymorphism")
	sa := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}
	p2 := person{
		first: "Jenny",
	}
	sa.speak()
	p2.speak()

	saySomething(sa)
	saySomething(p2)
}

func saySomething(h human) {
	fmt.Printf("Human of type %T is saying: ", h)
	h.speak()
}

func class139() {
	fmt.Println("\nClass 139 - Exploring the stringer interface")
	b := book{
		title: "Refactoring",
	}
	var c count = 42

	fmt.Println(b)
	fmt.Println(c)
}

func class140() {
	fmt.Println("\nClass 140 - Expanding on the stringer interface - wrapper func for logging")
	b := book{
		title: "Refactoring",
	}
	var c count = 42

	log.Println(b)
	log.Println(c)

	logExtraInfo(b)
	logExtraInfo(c)
}

func logExtraInfo(s fmt.Stringer) {
	log.Println("LOG FROM 140", s.String())
}

func class141() {
	fmt.Println("\nClass 141 - Writer interface & writing to a file")
	f, err := os.Create("sec-19/output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello Gophers!")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
