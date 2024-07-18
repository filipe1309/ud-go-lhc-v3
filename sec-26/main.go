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
	class202()
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

func class202() {
	fmt.Println("\nClass 202 - Race Condition")

	// runtime.GOMAXPROCS(100)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines before for\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg2 sync.WaitGroup
	wg2.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg2.Done()
			// Simulate race condition
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
		}()
		fmt.Println("GoRoutines in for\t", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines after for\t", runtime.NumGoroutine())
	wg2.Wait()
	fmt.Println("GoRoutines after wg2.Wait()\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
