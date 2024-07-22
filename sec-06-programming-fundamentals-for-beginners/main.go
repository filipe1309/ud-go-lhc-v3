package main

import (
	"fmt"
	"github.com/filipe1309/my-go-library"
	"github.com/filipe1309/ud-go-lhc-v3/sec-06-programming-fundamentals-for-beginners/otherpackage"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 6 - Programming fundamentals for beginners")
}

func main() {
	fmt.Println("Main function")
	otherpackage.ExportedFunction()
	mygolibrary.HelloFromL1v120()
}
