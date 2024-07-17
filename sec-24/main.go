package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
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
	class192()
	class193()
	class194()
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

func class192() {
	fmt.Println("\nClass 192 - Sort")
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("--------------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type Person3 struct {
	first string
	age   int
}

type ByAge []Person3

// Implementing sort.Interface, that is used by sort.Sort
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByName []Person3

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].first < a[j].first }

func class193() {
	fmt.Println("\nClass 193 - Sort custom")
	p1 := Person3{"James", 32}
	p2 := Person3{"Moneypenny", 27}
	p3 := Person3{"Q", 64}
	p4 := Person3{"M", 56}

	people := []Person3{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println("--------------")

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}

func class194() {
	fmt.Println("\nClass 194 - bcrypt")
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Println(string(bs))

	loginPassword := `password1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're logged in")
}
