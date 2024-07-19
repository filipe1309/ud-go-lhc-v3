package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 29 - Exercises Ninja Level 10")
}

func main() {
	fmt.Println("Main function")
	class222()
	class223()
}

func class222() {
	fmt.Println("\nClass 222 - Hands-on exercise #1")

	c := make(chan int)
	// c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func class223() {
	fmt.Println("\nClass 223 - Hands-on exercise #2")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Println("------")
	fmt.Printf("cs\t%T\n", cs)
}
