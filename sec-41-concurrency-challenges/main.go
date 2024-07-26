package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 41 - Concurrency Challenges")
}

func main() {
	fmt.Println("Main function")
	// class281()
	// class283()
	// class283V2()
	class283V3()
	class285()
	class287()
}

var workerID int
var publisherID int

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d:\tpushing data %d\n", thisID, dataID)
		data := fmt.Sprintf("Data %d from publisher %d", dataID, thisID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("worker %d:\twaiting for input...\n", thisID)
		input := <-in
		fmt.Printf("worker %d:\tinput is '%s'\n", thisID, input)
	}
}

func class281() {
	fmt.Println("\nClass 281 - Fan Out / Fan In - Challenge")

	input := make(chan string)

	// Fan Out
	go workerProcess(input)
	go workerProcess(input)
	go workerProcess(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(time.Millisecond)
}

var wg sync.WaitGroup

func class283Gen(n int) <-chan int {
	out := make(chan int)
	go func() {
		// insert n number in out channel
		for i := 1; i <= n; i++ {
			out <- (i % 10) // to avoid big numbers
		}
		close(out)
	}()
	return out
}

func class283FactorialRoutine(in <-chan int) chan int {
	out := make(chan int)
	for v := range in {
		go func(value int) {
			out <- class283Factorial(value)
			wg.Done()
		}(v)
	}
	return out
}

func class283Factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func class283() {
	fmt.Println("\nClass 283 - Fan Out / Fan In - Challenge: Factorial")

	n := 100
	wg.Add(n)
	in := class283Gen(n)
	f := class283FactorialRoutine(in)

	go func() {
		fmt.Println("#goroutine:", runtime.NumGoroutine())
		wg.Wait()
		close(f)
	}()

	for v := range f {
		fmt.Println(v)
	}
}

func class283V2FactorialRoutine(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- class283Factorial(v)
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg2 sync.WaitGroup
	wg2.Add(len(cs))

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg2.Done()
	}

	// launch output goroutines
	for _, c := range cs {
		go output(c)
	}

	go func() {
		fmt.Println("#goroutine:", runtime.NumGoroutine())
		wg2.Wait()
		close(out)
	}()

	return out
}

func class283V2() {
	fmt.Println("\nClass 283 v2 - Fan Out / Fan In - Challenge: Factorial")

	n := 10
	in := class283Gen(n)

	// Fan Out
	// launch 10 goroutines
	c0 := class283V2FactorialRoutine(in)
	c1 := class283V2FactorialRoutine(in)
	c2 := class283V2FactorialRoutine(in)
	c3 := class283V2FactorialRoutine(in)
	c4 := class283V2FactorialRoutine(in)
	c5 := class283V2FactorialRoutine(in)
	c6 := class283V2FactorialRoutine(in)
	c7 := class283V2FactorialRoutine(in)
	c8 := class283V2FactorialRoutine(in)
	c9 := class283V2FactorialRoutine(in)

	fmt.Println("#goroutine:", runtime.NumGoroutine())

	// Fan In
	for v := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println(v)
	}
}

func class283V3() {
	fmt.Println("\nClass 283 v3 - Fan Out / Fan In - Challenge: Factorial")

	n := 100
	in := class283Gen(n)

	// Fan Out
	chans := make([]<-chan int, n)
	// launch n goroutines
	for i := 0; i < n; i++ {
		ch := class283V2FactorialRoutine(in)
		chans[i] = ch
	}

	fmt.Println("#goroutine:", runtime.NumGoroutine())

	// Fan In
	for v := range merge(chans...) {
		fmt.Println(v)
	}
}

func class285FanOut(in <-chan int, n int) []<-chan int {
	// xc := make([]<-chan int, n) // This will create a deadlock
	var xc []<-chan int
	for i := 0; i < n; i++ {
		xc = append(xc, class283V2FactorialRoutine(in))
	}
	return xc
}

func class285() {
	fmt.Println("\nClass 285 - Deadlock Challenge")

	in := class283Gen(10)

	// Fan Out
	xc := class285FanOut(in, 10)

	// Fan In
	for v := range merge(xc...) {
		fmt.Println(v)
	}
}

var counter int64

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counter, 1)
		fmt.Println("Process", s, "printing:", i)
	}
	wg.Done()
}

func class287() {
	fmt.Println("\nClass 287 - Incrementor Challenge Revisited")

	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", counter) // 40
}
