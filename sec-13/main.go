package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 13 - Grouping data values - array & slice")
}

func main() {
	fmt.Println("Main function")
	class95()
	class96()
	class98()
	class99()
	class100()
	class101()
}

func class95() {
	fmt.Println("\nClass 95 - Array - an introduction to arrays")
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Printf("%v of type %T\n", a, a)

	b := [...]string{"James", "Bond", "Shaken, not stirred"}
	fmt.Printf("%v of type %T\n", b, b)
	// b = a // cannot use a (type [3]int) as type [3]string in assignment

	var c [2]int
	c[0] = 42
	c[1] = 43
	fmt.Printf("%v of type %T\n", c, c)

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}
}

func class96() {
	fmt.Println("\nClass 96 - Hands-on exercise #40")
	as := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Printf("length: %v\n", len(as))
	fmt.Printf("type: %T\n", as)
	fmt.Printf("array: %v\n", as)
}

func class98() {
	fmt.Println("\nClass 98 - Hands-on exercise #41")
	// slice literal
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Printf("length: %v\n", len(xs))
	fmt.Printf("type: %T\n", xs)
	fmt.Printf("slice: %v\n", xs)

	// slice
	for i, v := range xs {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
}

func class99() {
	fmt.Println("\nClass 99 - Slice - for range & access values by index position")
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Printf("length: %v\n", len(xs))
	fmt.Printf("type: %T\n", xs)
	fmt.Printf("slice: %v\n", xs)

	// blank identifier _ to not use a returned value
	for _, v := range xs {
		fmt.Printf("value: %v\n", v)
	}

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	// fmt.Println(xs[3]) // index out of range

	for i := 0; i < len(xs); i++ {
		fmt.Printf("index: %v, value: %v\n", i, xs[i])
	}
}

func class100() {
	fmt.Println("\nClass 100 - Slice - append to a slice")
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	xi = append(xi, 45, 46, 47)
	fmt.Println(xi)
}

func class101() {
	fmt.Println("\nClass 101 - Slice - slicing a slice")
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------------")
	
	fmt.Println("[inclusive:exclusive]")
	fmt.Printf("xi[0:4] - %#v\n", xi[0:4])
	fmt.Println("-------------------")
	
	fmt.Println("[:exclusive] - xi[:7] is the same as xi[0:7]")
	fmt.Printf("xi[:7] - %#v\n", xi[:7])
	fmt.Println("-------------------")

	fmt.Println("[inclusive:] - xi[4:] is the same as xi[4:len(xi)]")
	fmt.Printf("xi[4:] - %#v\n", xi[4:])
	fmt.Println("-------------------")

	fmt.Println("[:] - xi[:] is the same as xi[0:len(xi)]")
	fmt.Printf("xi[:] - %#v\n", xi[:])
	fmt.Println("-------------------")
}
