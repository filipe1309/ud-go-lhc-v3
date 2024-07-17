package main

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 24 - Application")
}

func main() {
	fmt.Println("Main function")
	class189()
}

type person struct {
	First string
	Last  string
	Age   int
}

func class189() {
	fmt.Println("\nClass 189 - JSON Marshal")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	xb, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}
