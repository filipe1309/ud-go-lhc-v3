package main

import (
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 25 - Exercises - Ninja Level 8")
}

func main() {
	fmt.Println("Main function")
	class195()
}

type user struct {
	First string
	Age   int
}

func class195() {
	fmt.Println("\nClass 195 - Hands-on exercise #1")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	users := []user{u1, u2}
	fmt.Println(users)

	xb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}
