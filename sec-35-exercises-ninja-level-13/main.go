package main

import (
	"fmt"

	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/dog"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 35 - Exercises - Ninja Level 13")
}

func main() {
	fmt.Println("Main function")
	class253()
}

type canine struct {
	name string
	age  int
}

func class253() {
	fmt.Println("\nClass 253 - Hands-on exercise #1")
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(10))
}
