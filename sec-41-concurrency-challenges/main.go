package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 41 - Concurrency Challenges")
}

func main() {
	fmt.Println("Main function")
	class281()
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
