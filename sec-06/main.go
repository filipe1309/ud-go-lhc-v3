package main

import (
	"fmt"
	"github.com/filipe1309/ud-go-lhc-v3/sec-06/otherpackage"
	"github.com/filipe1309/my-go-library"
)

func main() {
	fmt.Println("Main function")
	otherpackage.ExportedFunction()
	mygolibrary.HelloFromL1v120()
}
