package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 30 - Error Handling")
}

func main() {
	fmt.Println("Main function")
	class230()
	class230_3()
	class230_4()
}

func class230() {
	fmt.Println("\nClass 230 - Checking errors")

	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func class230_2() {
	fmt.Println("\nClass 230 - Checking errors 2")

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
}

func class230_3() {
	fmt.Println("\nClass 230 - Checking errors 3")

	f, err := os.Create("sec-30/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}

func class230_4() {
	fmt.Println("\nClass 230 - Checking errors 4")

	f, err := os.Open("sec-30/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
