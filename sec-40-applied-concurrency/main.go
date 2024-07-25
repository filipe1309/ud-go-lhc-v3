package main

import (
	"fmt"
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
