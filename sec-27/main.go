package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 27 - Exercises - Ninja Level 9")
}

func main() {
	fmt.Println("Main function")
	class207()
	class208()
	class209()
}

func class207() {
	fmt.Println("\nClass 207 - Hands-on exercise #1")
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("foo")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("bar")
	}()

	fmt.Println("GoRoutines after launch\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GoRoutines after wg.Wait()\t", runtime.NumGoroutine())
}

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func class208() {
	fmt.Println("\nClass 208 - Hands-on exercise #2")
	p1 := person{
		first: "James",
	}
	saySomething(&p1)
	// saySomething(p1) // This will not work, because p1 is not a pointer
	p1.speak() // This will work, because is not using the interface

	p2 := &person{
		first: "John",
	}
	saySomething(p2) // This will work, because p2 is a pointer
	p2.speak()
}

func class209() {
	fmt.Println("\nClass 209 - Hands-on exercise #3")
	// verify with: go run -race main.go

	var wg sync.WaitGroup
	const gs = 100
	counter := 0
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			// race condition
			v := counter
			runtime.Gosched()
			v++
			counter = v
		}()
	}

	wg.Wait()
	fmt.Println("Counter\t", counter)
}
