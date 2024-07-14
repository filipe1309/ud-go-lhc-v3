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
	class83()
	class84()
	class85()
	class86()
	class87()
	class89(class88())
	class90()
	class91()
	class92()
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

func class83() {
	fmt.Println("\nClass 83")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
}

func class84() {
	fmt.Println("\nClass 84")
	x := 0
	for {
		if x == 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}

func class85() {
	fmt.Println("\nClass 85")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func class86() {
	fmt.Println("\nClass 86")

	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("\tj: %v\n", j)
		}
	}
}

func class87() {
	fmt.Println("\nClass 87")
	x := []int{42, 43, 44, 45, 46, 47} // slice
	for i, v := range x {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
}

func class88() map[string]int {
	fmt.Println("\nClass 88")
	x := map[string]int{ // map
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range x {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	return x
}

func class89(x map[string]int) {
	fmt.Println("\nClass 89")
	age := x["James"]
	fmt.Printf("James age: %v\n", age)
	if v, ok := x["James"]; ok { // comma ok idiom
		fmt.Printf("James exists in the map, v: %v, ok: %v\n", v, ok)
	}

	age = x["Q"]
	fmt.Printf("Q age: %v\n", age)
	if v, ok := x["Q"]; !ok {
		fmt.Printf("Q does not exist in the map, v: %v, ok: %v\n", v, ok)
	}
}

func class90() {
	fmt.Println("\nClass 90")
	c := 1
	for i := 0; i < 100; i++ { // statement statement idiom
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("i: %v \t x: %v, \t #: %v\n", i, x, c)
			c++
		}
	}
}

func class91() {
	fmt.Println("\nClass 91")
	// fmt.Printf("true && true: \t%v\n", true && true)
	fmt.Printf("true && false: \t%v\n", true && false)
	// fmt.Printf("true || true: \t%v\n", true || true)
	fmt.Printf("true || false: \t%v\n", true || false)
	fmt.Printf("!true: \t\t%v\n", !true)
}

func class92() {
	fmt.Println("\nClass 92")
	// https://go.dev/play/p/DpZ_FLfn5s
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}

	// https://go.dev/play/p/IDnrJpE7vT
	x2 := "Moneypenny"

	if x2 == "Moneypenny" {
		fmt.Println(x2)
	} else if x2 == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x2)
	} else {
		fmt.Println("neither")
	}

	// https://go.dev/play/p/YpPgkWGqKY
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}

	// https://go.dev/play/p/Dj-yYGkAB4F
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("Living dangerously!")
	}
}
