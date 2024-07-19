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
