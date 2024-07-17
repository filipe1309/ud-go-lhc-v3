package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 25 - Exercises - Ninja Level 8")
}

func main() {
	fmt.Println("Main function")
	class195()
	class196()
	class197()
	class198()
	class199()
}

type user struct {
	First string
	Age   int
}

func class195() {
	fmt.Println("\nClass 195 - Hands-on exercise #1")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	users := []user{u1, u2}
	fmt.Println(users)

	xb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xb))
}

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func class196() {
	fmt.Println("\nClass 196 - Hands-on exercise #2")
	peopleJson := `[{"First":"James","Last":"Bond","Age":32, "Sayings":["Shaken, not stirred", "Youth is no guarantee of innovation", "In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27, "Sayings":["James! It is so good to see you!", "Would you like me to take care of that for you, James?"]}]`
	fmt.Println(peopleJson)

	var people []person
	err := json.Unmarshal([]byte(peopleJson), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, p := range people {
		fmt.Println("Person", i)
		fmt.Println("\tFirst:", p.First)
		fmt.Println("\tLast:", p.Last)
		fmt.Println("\tAge:", p.Age)
		for j, s := range p.Sayings {
			fmt.Println("\tSaying", j, ":", s)
		}
	}
}

func class197() {
	fmt.Println("\nClass 197 - Hands-on exercise #3")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	users := []user{u1, u2}

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func class198() {
	fmt.Println("\nClass 198 - Hands-on exercise #4")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2}
	xs := []string{"James", "M", "Bond", "Shaken, not stirred", "Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type byFirst []person

func (a byFirst) Len() int           { return len(a) }
func (a byFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFirst) Less(i, j int) bool { return a[i].First < a[j].First }

type byLast []person

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func class199() {
	fmt.Println("\nClass 199 - Hands-on exercise #5")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Age:     32,
		Sayings: []string{"Shaken, not stirred", "Youth is no guarantee of innovation", "In his majesty's royal service"},
	}

	p2 := person{
		First:   "Miss",
		Last:    "Moneypenny",
		Age:     27,
		Sayings: []string{"James! It is so good to see you!", "Would you like me to take care of that for you, James?"},
	}
	p3 := person{
		First:   "Q",
		Last:    "A",
		Age:     64,
		Sayings: []string{"I can do more damage on my laptop before my first cup of Earl Grey", "Oh, grow up, 007!"},
	}
	p4 := person{
		First:   "M",
		Last:    "M",
		Age:     56,
		Sayings: []string{"Oh, James. You didn't.", "Dear God, what has James done now?"},
	}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	fmt.Println("---------Original---------")
	for i, p := range people {
		fmt.Println("Person", i)
		fmt.Println("\tFirst:", p.First)
		fmt.Println("\tLast:", p.Last)
		fmt.Println("\tAge:", p.Age)
		for j, s := range p.Sayings {
			fmt.Println("\tSaying", j, ":", s)
		}
	}

	fmt.Println("---------By First---------")
	sort.Sort(byFirst(people))
	fmt.Println(people)
	for i, p := range people {
		fmt.Println("Person", i)
		fmt.Println("\tFirst:", p.First)
		fmt.Println("\tLast:", p.Last)
		fmt.Println("\tAge:", p.Age)
		sort.Strings(p.Sayings)
		for j, s := range p.Sayings {
			fmt.Println("\tSaying", j, ":", s)
		}
	}

	fmt.Println("---------By Last---------")
	sort.Sort(byLast(people))
	for i, p := range people {
		fmt.Println("Person", i)
		fmt.Println("\tFirst:", p.First)
		fmt.Println("\tLast:", p.Last)
		fmt.Println("\tAge:", p.Age)
		sort.Strings(p.Sayings)
		for j, s := range p.Sayings {
			fmt.Println("\tSaying", j, ":", s)
		}
	}
}
