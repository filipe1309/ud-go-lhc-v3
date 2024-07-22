package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 30 - Error Handling")
}

func main() {
	fmt.Println("Main function")
	// class230()
	// class230_3()
	// class230_4()
	// class231()
	// class232()
	class233()
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

func foo() {
	fmt.Println("foo")
}

func class231() {
	fmt.Println("\nClass 231 - Printing and logging")

	defer foo()

	// Config log to write to a file
	f, err := os.Create("sec-30/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	_, err = os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("err happened", err)
		// log.Panicln(err) // = Println() + panic()
		panic(err)
		// log.Fatalln(err) // = Println() + os.Exit(1)
	}

	fmt.Println("Check the log.txt file in the directory")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f()", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		// panic(fmt.Sprint("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func class232() {
	fmt.Println("\nClass 232 - Recover")

	f()
	fmt.Println("Returned normally from f()")
}

var ErrorNorgateMath = errors.New("norgate math: square root of negative number")

type norgateMathError struct {
	lat, long string
	err       error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, ErrorNorgateMath
		// return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f) // = errors.New(Sprintf())
		name := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", name}
	}
	return 42, nil
}

func class233() {
	fmt.Println("\nClass 233 - Errors with info")

	fmt.Printf("%T\n", ErrorNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
