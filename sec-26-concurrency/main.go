package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 26 - Concurrency")
}

func main() {
	fmt.Println("Main function")
	class201()
	class204()
	class205()
	class206()
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

func class204() {
	fmt.Println("\nClass 204 - Race Condition")

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

func class205() {
	fmt.Println("\nClass 205 - Mutex")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines before for\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg3 sync.WaitGroup
	wg3.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			defer wg3.Done()
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
		}()
		fmt.Println("GoRoutines in for\t", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines after for\t", runtime.NumGoroutine())
	wg3.Wait()
	fmt.Println("GoRoutines after wg3.Wait()\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}

func class206() {
	fmt.Println("\nClass 206 - Atomic")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines before for\t", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var wg4 sync.WaitGroup
	wg4.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg4.Done()
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter in goroutine\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
		}()
		fmt.Println("GoRoutines in for\t", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines after for\t", runtime.NumGoroutine())
	wg4.Wait()
	fmt.Println("GoRoutines after wg4.Wait()\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
