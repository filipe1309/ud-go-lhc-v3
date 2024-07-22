package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 14 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class108()
	class109()
	class110()
	class111()
	class112()
	class113()
	class114()
}

func class108() {
	fmt.Println("\nClass 108 - Hands-on exercise #42")
	// composite literal
	arr := [5]int{}
	for i := 0; i < 5; i++ {
		arr[i] = i * 2
	}

	for i, v := range arr {
		fmt.Printf("index: %v, value: %v, type: %T\n", i, v, v)
	}
}

func class109() {
	fmt.Println("\nClass 109 - Hands-on exercise #43")
	// composite literal
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range slice {
		fmt.Printf("index: %v, value: %v, type: %T\n", i, v, v)
	}
}

func class110() {
	fmt.Println("\nClass 110 - Hands-on exercise #44")
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// slice the slice
	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])
}

func class111() {
	fmt.Println("\nClass 111 - Hands-on exercise #45")
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slice)
	slice = append(slice, 52)
	fmt.Println(slice)
	slice = append(slice, 53, 54, 55)
	fmt.Println(slice)
	y := []int{56, 57, 58, 59, 60}
	slice = append(slice, y...)
	fmt.Println(slice)
}

func class112() {
	fmt.Println("\nClass 112 - Hands-on exercise #46")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// delete 45, 46, 47
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func class113() {
	fmt.Println("\nClass 113 - Hands-on exercise #47")
	states := make([]string, 0, 50)
	states = append(states,
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii",
		"Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota",
		"Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina",
		"North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah",
		"Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")

	fmt.Printf("lenght: %v\n", len(states))
	fmt.Printf("capacity: %v\n", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("%v: %v\n", i, states[i])
	}
}

func class114() {
	fmt.Println("\nClass 114 - Hands-on exercise #48")
	// a slice of a slice of string
	xxs := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."},
	}

	for i, xs := range xxs {
		fmt.Printf("record: %v\n", i)
		for j, v := range xs {
			fmt.Printf("\tindex	: %v\t value: %v\n", j, v)
		}
	}
}
