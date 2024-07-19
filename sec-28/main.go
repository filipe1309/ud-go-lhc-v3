package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 28 - Channels")
}

func main() {
	fmt.Println("Main function")
	class213()
	class214()
}

func class213() {
	fmt.Println("\nClass 213 - Understanding channels")

	fmt.Println("Channel with goroutine")
	c := make(chan int)
	// c <- 42 // This will cause a deadlock, because there is no goroutine to receive the value
	go func() { // This will work, because there is a goroutine to receive the value
		c <- 42
	}()
	fmt.Println(<-c)

	fmt.Println("Buffered channel")
	c2 := make(chan int, 1) // Buffered channel
	c2 <- 42                // This will work, because there is a buffer to store the value
	// c2 <- 43 // This will cause a deadlock, because the buffer is full
	fmt.Println(<-c2)

	fmt.Println("Buffered channel with multiple values")
	c3 := make(chan int, 2) // Buffered channel
	c3 <- 42
	c3 <- 43
	fmt.Println(<-c3)
	fmt.Println(<-c3)
}

func class214() {
	fmt.Println("\nClass 214 - Directional channels")

	c := make(chan int)
	cr := make(<-chan int) // Receive (from channel)
	cs := make(chan<- int) // Send (to channel)

	// cr <- 42 // This will cause an error, because cr is a receive-only channel
	go func() {
		cs <- 42 // This will work, because cs is a send-only channel
	}()

	fmt.Println("Channel")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// assign channel to directional channel
	// cr = c // This will work, because c is a bidirectional channel
	// cs = c // This will work, because c is a bidirectional channel

	fmt.Println("General to specific converts")
	fmt.Printf("c\t%T (send/receive)\n", c)
	fmt.Printf("c\t%T (send)\n", (chan<- int)(c))
	fmt.Printf("c\t%T (receive)\n", (<-chan int)(c))
}
