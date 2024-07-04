package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 15 - Grouping data values - map")
}

func main() {
	fmt.Println("Main function")
	class115()
	class116()
	class117()
}

func class115() {
	fmt.Println("\nClass 115 - Map - introduction")
	am := map[string]int {
		"Todd": 42,
		"Henry": 16,
		"Padget": 14,
	}

	fmt.Printf("The age of Todd is %v\n", am["Todd"])
	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	fmt.Printf("len: %v\n\n", len(am))

	an := make(map[string]int)
	an["Filipe"] = 34
	an["Quiky"] = 1
	fmt.Printf("The age of Filipe is %v\n", an["Filipe"])
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Printf("len: %v\n", len(an))
}

func class116() {
	fmt.Println("\nClass 116 - Map - for range over a map")

	// for range over a map
	am := map[string]int {
		"Todd": 42,
		"Henry": 16,
		"Padget": 14,
	}

	for k, v := range am {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	for _, v := range am {
		fmt.Printf("Value: %v\n", v)
	}

	for k := range am {
		fmt.Printf("Key: %v\n", k)
	}

	// for range with a slice
	xi := []int{42, 43, 44}

	for i, v := range xi {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	for _, v := range xi {
		fmt.Printf("Value: %v\n", v)
	}

	for i := range xi {
		fmt.Printf("Index: %v\n", i)
	}
}

func class117() {
	fmt.Println("\nClass 117 - Map - delete element")

	am := map[string]int {
		"Todd": 42,
		"Henry": 16,
		"Padget": 14,
	}

	fmt.Println(am)
	delete(am, "Todd")
	delete(am, "NotExists") // no error, won't panic
	fmt.Println(am)
	fmt.Println(am["NotExists"]) // zero value, no error, won't panic
}
