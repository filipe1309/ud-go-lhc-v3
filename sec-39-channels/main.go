package main

import (
	"fmt"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 39 - Channels")
}

func main() {
	fmt.Println("Main function")
	class264()
	class265()
	class266()
}

func class264() {
	fmt.Println("\nClass 264 - Channels - Introduction")

	// Unbuffered channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// Send value to channel
			c <- i
		}
	}()

	go func() {
		for {
			// Receive value from channel, blocks until a value is received
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second) // Sleep for a second to allow the goroutines to finish
}

func class265() {
	fmt.Println("\nClass 265 - Range clause")

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending value to channel:", i)
			// Send value to channel
			c <- i
		}
		close(c) // Close the channel, no more values can be sent, but can still receive values
	}()

	// Receive values from channel
	for v := range c {
		fmt.Println("Received value from channel:", v)
		fmt.Println(v)
	}
}

func class266() {
	fmt.Println("\nClass 266 - N-to-1")

	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending value to channel:", i, "from goroutine 1")
			// Send value to channel
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending value to channel:", i, "from goroutine 2")
			// Send value to channel
			c <- i
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("Waiting for goroutines to finish from goroutine 3")
		wg.Wait()
		fmt.Println("Closing channel from goroutine 3")
		close(c)
	}()

	// Receive values from channel
	for v := range c {
		fmt.Println("Received value from channel:", v)
	}
}
