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
