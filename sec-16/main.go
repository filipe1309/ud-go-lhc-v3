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
	things := class120()
	printMap(things)
	class121(things)
	printMap(things)
}

func printMap(things map[string][]string) {
	for k, v := range things {
		fmt.Printf("Key: %v\n", k)
		for i, v2 := range v {
			fmt.Printf("\tIndex: %v, Value: %v\n", i, v2)
		}
	}
}

func class120() map[string][]string {
	fmt.Println("\nClass 120 - Hands-on exercise #49 - map[string][]string")
	things := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	fmt.Printf("%#v\n", things)
	return things
}

func class121(things map[string][]string) {
	fmt.Println("\nClass 121 - Hands-on exercise #50 - add a record")
	things["fleming_ian"] = []string{"streaks", "cigars", "espionage"}
}
