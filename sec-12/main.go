package main

import (
	"fmt"
	"math/rand"
)

// Class 78
// https://go.dev/doc/effective_go#init
func init() {
	fmt.Println("Init function")
	class78()
}

func main() {
	fmt.Println("Main function")
	class76()
	class80(class79())
	class81()
	class82()
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

	switch {
	case x <= 100:
		fmt.Println("between 0 and 100")
	case x <= 200:
		fmt.Println("between 101 and 200")
	default:
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

func class78() {
	fmt.Println("\nClass 78")
	fmt.Println("This is where initialization for my program occurs")
}

func class79() (x, y int) {
	fmt.Println("\nClass 79")
	x = rand.Intn(10)
	y = rand.Intn(10)
	fmt.Printf("x: %v, y: %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous cases were met")
	}

	return
}

func class80(x, y int) {
	fmt.Println("\nClass 80")
	class80Switch(x, y)
	for i := 0; i < 100; i++ {
		x = rand.Intn(10)
		y = rand.Intn(10)
		fmt.Printf("x: %v, y: %v\n", x, y)
		class80Switch(x, y)
	}
}

func class80Switch(x, y int) {
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}

func class81() {
	fmt.Println("\nClass 81")
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func class82() {
	fmt.Println("\nClass 82")
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Printf("%v: \t x is %v\n", i, x)
		case 1:
			fmt.Printf("%v: \t x is %v\n", i, x)
		case 2:
			fmt.Printf("%v: \t x is %v\n", i, x)
		case 3:
			fmt.Printf("%v: \t x is %v\n", i, x)
		case 4:
			fmt.Printf("%v: \t x is %v\n", i, x)
		}
	}
}
