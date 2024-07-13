package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 20 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class152()
}

func class152() {
	fmt.Println("\nClass 152 - Hands-on exercise #57 - func concepts")
	x := sumNamedReturn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
	x = sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sumNamedReturn(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

func sum(ii []int) (int) {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}
