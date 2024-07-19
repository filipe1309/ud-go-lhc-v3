package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 29 - Exercises Ninja Level 10")
}

func main() {
	fmt.Println("Main function")
	class222()
	class223()
	class224()
	class225()
	class226()
	class227()
}

func class222() {
	fmt.Println("\nClass 222 - Hands-on exercise #1")

	c := make(chan int)
	// c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func class223() {
	fmt.Println("\nClass 223 - Hands-on exercise #2")

	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Println("------")
	fmt.Printf("cs\t%T\n", cs)
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func class224() {
	fmt.Println("\nClass 224 - Hands-on exercise #3")

	c := gen()
	receive(c)

	fmt.Println("About to exit")
}

func gen225(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive225(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("From c:", v)
		case v := <-q:
			fmt.Println("From q:", v)
			return
		}
	}
}

func class225() {
	fmt.Println("\nClass 225 - Hands-on exercise #4")

	q := make(chan int)
	c := gen225(q)

	receive225(c, q)

	fmt.Println("About to exit")
}

func class226() {
	fmt.Println("\nClass 226 - Hands-on exercise #5")

	c := make(chan int, 1)

	c <- 42

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func class227() {
	fmt.Println("\nClass 227 - Hands-on exercise #6")

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}
