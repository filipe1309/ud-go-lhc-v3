package main

import (
	"fmt"
	"github.com/filipe1309/ud-go-lhc-v3/sec-33-exercises-ninja-level-12/dog"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 33 - Exercises - Ninja Level 12")
}

func main() {
	fmt.Println("Main function")
	class244()
}

type canine struct {
	name string
	age  int
}

func class244() {
	fmt.Println("\nClass 244 - Hands-on exercise #1")

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)

	// go doc sec-33-exercises-ninja-level-12/dog
	// go doc sec-33-exercises-ninja-level-12/dog.Years
}
