package main

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 20 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class152()
	class153()
	class154()
	class155()
	class156()
	class157()
	class158()
	class159()
}

func class152() {
	fmt.Println("\nClass 152 - Hands-on exercise #57 - func concepts")
	x := sumNamedReturn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
	x = sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sumNamedReturn(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func class153() {
	fmt.Println("\nClass 153 - Hands-on exercise #58 - basic funcs")
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 21, "James Bond"
}

func class154() {
	fmt.Println("\nClass 154 - Hands-on exercise #59 - variadic func")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo2(xi...))
	fmt.Println(bar2(xi))
}

func foo2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func class155() {
	fmt.Println("\nClass 155 - Hands-on exercise #60 - defer func")
	fmt.Println("Before defer funcs")
	// LIFO
	defer fmt.Println("Deferred func 1")
	defer fmt.Println("Deferred func 2")
	defer fmt.Println("Deferred func 3")
	fmt.Println("After defer funcs")
}

func class156() {
	fmt.Println("\nClass 156 - Hands-on exercise #61 - method")
	p1 := person{
		first: "John",
		age:   42,
	}
	p1.speak()
}

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Name: %v, Age: %v\n", p.first, p.age)
}

func class157() {
	fmt.Println("\nClass 157 - Hands-on exercise #62 - interfaces")
	s := square{
		length: 5,
		width:  8,
	}
	c := circle{
		radius: 4,
	}

	info(s)
	info(c)
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%T - area: %v\n", s, s.area())
}

func class158() {
	fmt.Println("\nClass 158 - Hands-on exercise #63 - tests in go #1")
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}

func class159() {
	fmt.Println("\nClass 159 - Hands-on exercise #64 - tests in go #2 - unit tests")
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}
