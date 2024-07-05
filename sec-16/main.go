package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 16 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class120()
}

func class120() {
	fmt.Println("\nClass 120 - Hands-on exercise #49 - map[string][]string")
	things := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	

	for k, v := range things {
		fmt.Printf("Key: %v\n", k)
		for i, v2 := range v {
			fmt.Printf("\tIndex: %v, Value: %v\n", i, v2)
		}
	}
}
