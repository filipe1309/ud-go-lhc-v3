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
	class236()
	class237()
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

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	// fmt.Println("foo ran -", e)
	fmt.Println("foo ran -", e.(customErr).info) // e.(customErr) = Assertion
}

func class236() {
	fmt.Println("\nClass 236 - Hands-on exercise #3")

	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("a sqrt error occurred: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e :=  errors.New("more coffee needed")
		e := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

func class237() {
	fmt.Println("\nClass 237 - Hands-on exercise #4")

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
