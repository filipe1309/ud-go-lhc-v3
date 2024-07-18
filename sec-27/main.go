package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 27 - Exercises - Ninja Level 9")
}

func main() {
	fmt.Println("Main function")
	class207()
}

func class207() {
	fmt.Println("\nClass 207 - Hands-on exercise #1")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("foo")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("bar")
	}()

	fmt.Println("GoRoutines after launch\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GoRoutines after wg.Wait()\t", runtime.NumGoroutine())
}
