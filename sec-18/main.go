package main

import (
	"fmt"
)

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make string
	model string
	doors int
	color string
}

type person struct {
	first string
	last string
	flavors []string
}

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 18 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class128()
	class129()
	class130()
	class131()
}

func class128() {
	fmt.Println("\nClass 128 - Hands-on exercise #53 - struct with slice")

	p1 := person{
		first: "John",
		last: "Doe",
		flavors: []string{"vanilla", "cream"},
	}

	p2 := person{
		first: "Elvis",
		last: "Presley",
		flavors: []string{"chocolate", "strwberry", "lemon"},
	}

	fmt.Println(p1)

	for _, v := range p1.flavors {
		fmt.Println(v)
	}

	fmt.Println(p2)

	for _, v := range p2.flavors {
		fmt.Println(v)
	}
}

func class129() {
	fmt.Println("\nClass 129 - Hands-on exercise #54 - map struct")

	p1 := person{
		first: "John",
		last: "Doe",
		flavors: []string{"vanilla", "cream"},
	}

	p2 := person{
		first: "Elvis",
		last: "Presley",
		flavors: []string{"chocolate", "strwberry", "lemon"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Printf("Key: %v\n", k)
		for _, v2 := range v.flavors {
			fmt.Printf("\t Value: %v\n", v2)
		}
	}
}


func class130() {
	fmt.Println("\nClass 130 - Hands-on exercise #55 - embed struct")

	v1 := vehicle{
		engine: engine{
			electric: true,
		},
		make: "BYD",
		model: "Dolphin",
		doors: 2,
		color: "red",
	}

	v2 := vehicle{
		engine: engine{
			electric: false,
		},
		make: "Ford",
		model: "Ka",
		doors: 4,
		color: "black",
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Printf("Engine: %v, Make: %v, Model: %v, Doors: %v, Color: %v\n", v1.electric, v1.make, v1.model, v1.doors, v1.color)
	fmt.Printf("Engine: %v, Make: %v, Model: %v, Doors: %v, Color: %v\n", v2.engine.electric, v2.make, v2.model, v2.doors, v2.color)
}

func class131() {
	fmt.Println("\nClass 131 - Hands-on exercise #55 - anonymous struct")

	p1 := struct {
		fisrt string
		friends map[string]int
		favDrinks []string
	}{
		fisrt: "John",
		friends: map[string]int{"Elvis": 38, "James": 42},
		favDrinks: []string{"Martini", "Vodka"},
	}

	fmt.Println(p1)

	for k, v := range p1.friends {
		// fmt.Printf("Key: %v \t Value: %v\n", k, v)
		fmt.Println("Key:", k, "\t Value: ", v)
	}

	for i, v := range p1.favDrinks {
		fmt.Printf("Index: %v \t Value: %v\n", i, v)
	}
}

