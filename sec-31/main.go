package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 31 - Exercises - Ninja Level 11")
}

func main() {
	fmt.Println("Main function")
	class234()
	class235()
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func class234() {
	fmt.Println("\nClass 234 - Hands-on exercise #1")

	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never",
		},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
		// return []byte{}, errors.New(fmt.Sprintf("there was an error in toJSON: %v", err))
	}
	return bs, nil

}

func class235() {
	fmt.Println("\nClass 235 - Hands-on exercise #2")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Sayings: []string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never",
		},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}
	fmt.Println(string(bs))
}
