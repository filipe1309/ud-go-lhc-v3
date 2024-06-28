package otherpackage

import (
	"fmt"
)

// ExportedFunction is a function that can be called from other packages because it starts with an uppercase letter
func ExportedFunction() {
	fmt.Println("Exported function")
	unexportedFunction()
}

// unexportedFunction is a function that can't be called from other packages because it starts with a lowercase letter
func unexportedFunction() {
	fmt.Println("Unexported function")
}
