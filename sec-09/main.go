package main

import (
	"fmt"
)

var myVarAtPackageLevel = "My var at package level"
const myConstAtPackageLevel = "My const at package level"

func main() {
	fmt.Println("Main function")
	fmt.Println("Class 54")
	fmt.Printf("myVarAtPackageLevel: %v\n", myVarAtPackageLevel)
	fmt.Printf("myConstAtPackageLevel: %v\n", myConstAtPackageLevel)

	myVarAtBlockLevel := "My var at block level"
	const myConstAtBlockLevel = "My const at block level"
	fmt.Printf("myVarAtBlockLevel: %v\n", myVarAtBlockLevel)
	fmt.Printf("myConstAtBlockLevel: %v\n", myConstAtBlockLevel)
}

