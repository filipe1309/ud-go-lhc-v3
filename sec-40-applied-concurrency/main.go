package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 40 - Applied Concurrency")
}

func main() {
	fmt.Println("Main function")
	class272()
	class273()
	class273_2()
	class274()
	class275()
	class276()
	class279()
	class280()
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for v := range c {
			sum += v
		}
		out <- sum
		close(out)
	}()
	return out
}

func class272() {
	fmt.Println("\nClass 272 - Incrementor With Channels")

	c1 := incrementor("foo:")
	c2 := incrementor("bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}

func class273() {
	fmt.Println("\nClass 273 - Deadlock Challenge")

	// c := make(chan int)
	// 	c <- 1 // deadlock because there is no goroutine to receive the value, so it blocks. Or use a buffered channel
	// fmt.Println(<-c)

	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

func class273_2() {
	fmt.Println("\nClass 273 - Deadlock Challenge 2")

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // for range loop
	}()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }

	for v := range c {
		fmt.Println(v)
	}
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func factorialGoroutAndChan(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func class274() {
	fmt.Println("\nClass 274 - Factorial Challenge")

	f := factorial(4)
	fmt.Println("Total:", f)

	// This version o factorial with goroutines and channel will have a better performance
	// if you need to calculate thousands os factorial operations, because you will be
	// able to use all CPU cores of your machine to do that
	c := factorialGoroutAndChan(4)
	for v := range c {
		fmt.Println("Total 2:", v)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func sqrt(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func class275() {
	fmt.Println("\nClass 275 - Pipeline Pattern")

	// for v := range sqrt(sqrt(gen(2, 3))) {
	// 	fmt.Println(v) // 16, 81
	// }

	// Set up the pipeline
	c := gen(2, 3)
	out := sqrt(c)

	// Consume the output
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen276() <-chan int {
	out := make(chan int)
	go func() {
		// insert 100 number in out channel
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial276(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- factorial(v)
		}
		close(out)
	}()
	return out
}

func class276() {
	fmt.Println("\nClass 276 - Factorial Challenge Redux")

	in := gen276()
	f := factorial276(in)

	for v := range f {
		fmt.Println(v)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func class279() {
	fmt.Println("\nClass 279 - Fan In Pattern")
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func mergeFanIn(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, v := range cs {
		go func(ch <-chan int) {
			for c := range ch {
				out <- c
			}
			wg.Done()
		}(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func class280() {
	fmt.Println("\nClass 280 - Fan Out / Fan In - Example")

	// Fan Out
	// Distribute the sqrt work across two goroutines that both read from in.
	in := gen(2, 3)
	c1 := sqrt(in)
	c2 := sqrt(in)

	// Fan In
	for v := range mergeFanIn(c1, c2) {
		fmt.Println(v) // 4, 9 or 9, 4
	}
}
