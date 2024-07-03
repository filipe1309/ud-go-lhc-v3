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
