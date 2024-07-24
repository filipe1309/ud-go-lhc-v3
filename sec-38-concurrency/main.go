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
	class260()
	class261()
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

var counter int

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(3 * time.Millisecond))
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run -race sec-38-concurrency/main.go
func class260() {
	fmt.Println("\nClass 260 - Race Conditions")
	wg.Add(2)
	go incrementor("foo:")
	go incrementor("bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

var mu sync.Mutex

func incrementorWithMutex(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(20 * time.Millisecond))
		mu.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mu.Unlock()
	}
	wg.Done()
}

func class261() {
	fmt.Println("\nClass 261 - Mutex")
	counter = 0
	wg.Add(2)
	go incrementorWithMutex("foo:")
	go incrementorWithMutex("bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
