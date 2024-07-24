package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 39 - Channels")
}

func main() {
	fmt.Println("Main function")
	class264()
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

	time.Sleep(time.Second)
}
