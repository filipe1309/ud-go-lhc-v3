package main

import (
	"fmt"
	"log"
	"math"
)

func init() {
	fmt.Println("Init function")
	fmt.Println("Section 20 - Hands-on exercises")
}

func main() {
	fmt.Println("Main function")
	class152()
	class153()
	class154()
	class155()
	class156()
	class157()
	class158()
	class159()
	class160()
	class162()
	class163()
	class164()
	class165()
	class166()
}

func class152() {
	fmt.Println("\nClass 152 - Hands-on exercise #57 - func concepts")
	x := sumNamedReturn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
	x = sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sumNamedReturn(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

func sum(ii []int) int {
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

func class153() {
	fmt.Println("\nClass 153 - Hands-on exercise #58 - basic funcs")
	x := foo()
	fmt.Println(x)
	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 21, "James Bond"
}

func class154() {
	fmt.Println("\nClass 154 - Hands-on exercise #59 - variadic func")
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo2(xi...))
	fmt.Println(bar2(xi))
}

func foo2(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func class155() {
	fmt.Println("\nClass 155 - Hands-on exercise #60 - defer func")
	fmt.Println("Before defer funcs")
	// LIFO
	defer fmt.Println("Deferred func 1")
	defer fmt.Println("Deferred func 2")
	defer fmt.Println("Deferred func 3")
	fmt.Println("After defer funcs")
}

func class156() {
	fmt.Println("\nClass 156 - Hands-on exercise #61 - method")
	p1 := person{
		first: "John",
		age:   42,
	}
	p1.speak()
}

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Name: %v, Age: %v\n", p.first, p.age)
}

func class157() {
	fmt.Println("\nClass 157 - Hands-on exercise #62 - interfaces")
	s := square{
		length: 5,
		width:  8,
	}
	c := circle{
		radius: 4,
	}

	info(s)
	info(c)
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%T - area: %v\n", s, s.area())
}

func class158() {
	fmt.Println("\nClass 158 - Hands-on exercise #63 - tests in go #1")
	fmt.Println(add(3, 4))
}

func add(a, b int) int {
	return a + b
}

func class159() {
	fmt.Println("\nClass 159 - Hands-on exercise #64 - tests in go #2 - unit tests")
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func subtract(a, b int) int {
	return a - b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func class160() {
	fmt.Println("\nClass 160 - Hands-on exercise #65 - tests in go #3 - unit tests")
	fmt.Println(Paradise("Buenos Aires"))
}

// Paradise prints out the user's input of paradise as a setence
func Paradise(loc string) string {
	return fmt.Sprintf("My idea of paradise is %s", loc)
}


// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// MockDatastore is a temporary service that stores retrievable data.
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`.
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.DB` and it can also be of TYPE `Datastore`
// This means we can write functions to take TYPE `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
// https://pkg.go.dev/database/sql#Open
type Datastore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Service represents a service that stores retrievable data.
// We will attach methods to `Service` so that we can use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Service struct {
	ds Datastore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func class162() {
	fmt.Println("\nClass 162 - Hands-on exercise #67 - interfaces & mock testing a db")
	db := MockDatastore{
		Users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
}

func class163() {
	fmt.Println("\nClass 163 - Hands-on exercise #68 - anonymous func")
	// func(){}()
	func () {
		fmt.Println("My anonymous func")
	}()
}

func class164() {
	fmt.Println("\nClass 164 - Hands-on exercise #69 - func expression")
	x := func () {
		fmt.Println("My func expression")
	}
	x()
}

func class165() {
	fmt.Println("\nClass 165 - Hands-on exercise #70 - func return")
	f := myOuterFunc()
	fmt.Println(f())
}

func myOuterFunc() func () int {
	return func () int {
		return 42
	}
}

func class166() {
	fmt.Println("\nClass 166 - Hands-on exercise #71 - callback")
	fmt.Println(genSquareText(5, calcSquare))
}

func genSquareText(i int, f func(int) int) string {
	s := f(i)
	return fmt.Sprintf("the number %v squared is %v", i, s)
}

func calcSquare(i int) int {
	return i * i
}

