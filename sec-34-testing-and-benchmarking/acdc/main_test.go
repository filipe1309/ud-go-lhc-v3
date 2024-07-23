package acdc

import "fmt"

// Example tests
// go test -v sec-34-testing-and-benchmarking/acdc/*
func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}

// func TestSum(t *testing.T) {
// 	got := Sum(2, 3)
// 	want := 5
// 	if got != want {
// 		t.Error("Want", want, "got", got)
// 	}
// }
