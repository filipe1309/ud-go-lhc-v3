package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 24 - Application")
}

func main() {
	fmt.Println("Main function")
	class189()
	class190()
	class191()
}

type person struct {
	First string
	Last  string
	Age   int
}

func class189() {
	fmt.Println("\nClass 189 - JSON Marshal")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	xb, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}

type person2 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func class190() {
	fmt.Println("\nClass 190 - JSON Unmarshal")
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)
	fmt.Printf("%T - %v\n", s, s)
	fmt.Printf("%T - %v\n", bs, bs)

	var people []person2
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)
	for i, v := range people {
		fmt.Println("Person #", i)
		fmt.Println("\tFirst:", v.First)
		fmt.Println("\tLast:", v.Last)
		fmt.Println("\tAge:", v.Age)
	}
}

func class191() {
	fmt.Println("\nClass 191 - Writer interface")
	fmt.Println("Hello World!")
	fmt.Fprintln(os.Stdout, "Hello World!") // os.Stdout implements Writer interface
	io.WriteString(os.Stdout, "Hello World!\n")
}
