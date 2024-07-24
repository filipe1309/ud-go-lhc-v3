package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 38 - Concurrency")
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the cores - Parallelism
}

func main() {
	fmt.Println("Main function")
	class258()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func class258NoConcurrency() {
	foo()
	bar()
}

var wg sync.WaitGroup

func foo2() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar2() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

func class258WithConcurrency() {
	wg.Add(2)
	go foo2()
	go bar2()
	wg.Wait()
}

func class258() {
	fmt.Println("\nClass 258 - Concurrency & WaitGroup")

	fmt.Println("\nNo Concurrency")
	class258NoConcurrency()

	fmt.Println("\nWith Concurrency")
	class258WithConcurrency()
}
