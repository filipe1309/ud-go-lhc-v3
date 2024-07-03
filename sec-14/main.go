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
}

func class108() {
	fmt.Println("\nClass 108 - Hands-on exercise #42")
	// composite literal
	arr := [5]int{}
	for i := 0; i < 5; i++ {
		arr[i] = i*2
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
