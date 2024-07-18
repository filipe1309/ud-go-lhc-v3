package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 26 - Concurrency")
}

func main() {
	fmt.Println("Main function")
	class201()
}

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func class201() {
	fmt.Println("\nClass 201 - WaitGroup")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
}
