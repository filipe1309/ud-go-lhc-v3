package main

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
	"github.com/filipe1309/my-go-library"
)

var myVarAtPackageLevel = "My var at package level"

const myConstAtPackageLevel = "My const at package level"

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 9 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	fmt.Println("Class 54 - Hands-on exercise #14")
	fmt.Printf("myVarAtPackageLevel: %v\n", myVarAtPackageLevel)
	fmt.Printf("myConstAtPackageLevel: %v\n", myConstAtPackageLevel)

	myVarAtBlockLevel := "My var at block level"
	const myConstAtBlockLevel = "My const at block level"
	fmt.Printf("myVarAtBlockLevel: %v\n", myVarAtBlockLevel)
	fmt.Printf("myConstAtBlockLevel: %v\n", myConstAtBlockLevel)

	class59()
	class60()
}

func class59() {
	fmt.Println("\nClass 59 - Hands-on exercise #19")
	mygolibrary.HelloFromL1v120()
}

func class60() {
	fmt.Println("\nClass 60 - Hands-on exercise #20")
	fmt.Printf("puppy.Barks(): %v\n", puppy.Barks())
}
