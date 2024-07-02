package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Main function")
	class76()
}

func class76() {
	fmt.Println("\nClass 76")
	x := rand.Intn(251)
	fmt.Printf("x: %v\n", x)
	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}

	// rand.Intn(3) returns a random number between 0 and 2
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
	fmt.Println("rand.Intn(3):", rand.Intn(3))
}
