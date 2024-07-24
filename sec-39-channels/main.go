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
	class267()
	class268()
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

func class267() {
	fmt.Println("\nClass 267 - Semaphores - Part 1")

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending\t", i, "to channel (goroutine 1)")
			// Send value to channel
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending\t", i, "to channel (goroutine 2)")
			// Send value to channel
			c <- i
		}
		done <- true
	}()

	go func() {
		fmt.Println("--------------------------------")
		fmt.Println("Waiting for goroutines to finish (goroutine 3)")
		fmt.Println("--------------------------------")
		<-done
		fmt.Println("--------------------------------")
		fmt.Println("One goroutine finished")
		fmt.Println("--------------------------------")
		<-done
		fmt.Println("--------------------------------")
		fmt.Println("Two goroutines finished")
		fmt.Println("--------------------------------")
		fmt.Println("Closing channel (goroutine 3)")
		fmt.Println("--------------------------------")
		close(c)
	}()

	// Receive values from channel
	for v := range c {
		fmt.Println("Received from channel:\t", v)
	}
}

func printWithBorders(s string) {
	fmt.Println("--------------------------------")
	fmt.Println(s)
	fmt.Println("--------------------------------")
}

func class268() {
	fmt.Println("\nClass 268 - Semaphores - Part 2")

	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for i2 := 0; i2 < 10; i2++ {
				fmt.Println("Sending\t", i2, "to channel ( goroutine", i, ")")
				// Send value to channel
				c <- i2
			}
			done <- true
		}(i)
	}

	go func() {
		printWithBorders("Waiting for goroutines to finish (goroutine 3)")
		for i := 0; i < n; i++ {
			<-done
			printWithBorders(fmt.Sprintf("%d goroutine(s) finished (goroutine 3)", i+1))
		}
		printWithBorders("Closing channel (goroutine 3)")
		close(c)
	}()

	// Receive values from channel
	for v := range c {
		fmt.Println("Received from channel:\t", v)
	}
}
