package main

import (
	"encoding/json"
	"fmt"
	"os"
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
