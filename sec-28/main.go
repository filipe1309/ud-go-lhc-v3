package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 28 - Channels")
}

func main() {
	fmt.Println("Main function")
	class213()
	class214()
	class215()
	class216()
	class217()
	class218()
	class219()
	class220()
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

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}

func class215() {
	fmt.Println("\nClass 215 - Using channels")

	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("End of the function")
}

func class216() {
	fmt.Println("\nClass 216 - Range")

	c := make(chan int)

	// Send
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // Close the channel, so the receiver will know that there is no more value to receive
	}()

	// Receive
	// range wait for the channel to be closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("End of the function")
}

func sendTo(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}

func receiveFrom(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even:\t", v)
		case v := <-odd:
			fmt.Println("Odd:\t", v)
		case v := <-quit:
			fmt.Println("Quit:\t", v)
			return
		}
	}
}

func class217() {
	fmt.Println("\nClass 217 - Select")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go sendTo(even, odd, quit)

	// Receive
	receiveFrom(even, odd, quit)

	fmt.Println("End of the function")
}

func sendTo218(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receiveFrom218(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("Even:\t", v)
		case v := <-odd:
			fmt.Println("Odd:\t", v)
		case v, ok := <-quit:
			fmt.Println("Quit:\t", v)
			if !ok {
				fmt.Println("Quit channel closed")
				return
			}
		}
	}
}

func class218() {
	fmt.Println("\nClass 218 - Comma ok idiom")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Send
	go sendTo218(even, odd, quit)

	// Receive
	receiveFrom218(even, odd, quit)

	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)

	fmt.Println("End of the function")
}

func sendTo219(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

func receiveFrom219(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range even {
			fanin <- v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range odd {
			fanin <- v
		}
	}()

	wg.Wait()
	close(fanin)
}

func class219() {
	fmt.Println("\nClass 219 - Fan in")

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// Send
	go sendTo219(even, odd)

	// Receive
	go receiveFrom219(even, odd, fanin)

	for v := range fanin {
		fmt.Println("Fan in:\t", v)
	}

	fmt.Println("End of the function")
}

func populate(c1 chan int) {
	for i := 0; i < 5; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		// Fan out (multiple goroutines), passing the values to the goroutines
		go func(v2 int) {
			// Fan in, send the result to c2
			fmt.Printf("Go routine: %d - %d\n", runtime.NumGoroutine(), v2)
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}

func class220() {
	fmt.Println("\nClass 220 - Fan out")

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println("Fan out:\t", v)
	}

	fmt.Println("End of the function")
}
