package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 21 - Pointers")
}

func main() {
	fmt.Println("Main function")
	class169()
	class170()
	class171()
	class172()
	class173()
	class175()
}

func class169() {
	fmt.Println("\nClass 169 - What are pointers?")
	x := 42
	fmt.Println(x) // x value
	fmt.Println(&x) // x address
}

func class170() {
	fmt.Println("\nClass 170 - Seeing type & value for pointers?")
	x := 42
	fmt.Printf("x value: %v, x type: %T\n", &x, &x)

	y := "James Bond"
	fmt.Printf("y value: %v, y type: %T\n", &y, &y)
}

func class171() {
	fmt.Println("\nClass 171 - Defereferencing pointers")
	x := 42
	fmt.Printf("x value: %v, x type: %T\n", x, x)

	y := &x
	fmt.Printf("y value: %v, y type: %T\n", y, y)
	fmt.Printf("*y value: %v, *y type: %T\n", *y, *y)
	fmt.Printf("y address: %v, y type: %T\n", &y, &y)
	// fmt.Println(*&x) // dereferencing

	*y = 43
	fmt.Printf("x value: %v, x type: %T\n", x, x)
	fmt.Printf("*y value: %v, *y type: %T\n", *y, *y)
}

func class172() {
	fmt.Println("\nClass 172 - Pass by value, pointers / reference types, and mutability")
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	// slice is a reference type, it is a pointer to an array, so it is passed by reference
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	// map is a reference type, it is a pointer to a hash table, so it is passed by reference
	m := map[string]int{
		"James": 32,
		"Jenny": 27,
	}
	fmt.Println(m)
	mapDelta(m, "James")
	fmt.Println(m)
}

func intDelta(a *int) {
	*a = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(mm map[string]int, s string) {
	mm[s] = 42
}

func class173() {
	fmt.Println("\nClass 173 - Pointer & value semantics")
	a := 1
	fmt.Println(a) // 1
	fmt.Println(addOneValue(a)) // 2
	fmt.Println(a) // 1

	b := 1
	fmt.Println(b) // 1
	fmt.Println(addOnePointer(&b)) // 2
	fmt.Println(b) // 2
}

// value semantics
func addOneValue(a int) int {
	a++
	return a
}

// pointer semantics
func addOnePointer(a *int) int {
	*a++
	return *a
}

func class175() {
	fmt.Println("\nClass 175 - Pointers, values, the stack & the heap")
	a := 1
	fmt.Println(a) // "a escapes to heap", stays in stack, because it is not passed to a function by value

	b := 1
	fmt.Println(&b) // "moved to heap: b", because it is passed to a function by reference
}
