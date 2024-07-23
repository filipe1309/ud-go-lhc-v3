package mymath

import (
	"fmt"
	"testing"
)

// go test -bench Center github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/mymath
func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

// go test -v sec-35-exercises-ninja-level-13/mymath/*
func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 4
}

// go test -v sec-35-exercises-ninja-level-13/mymath/*
func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4.0},
		{[]int{10, 20, 40, 60, 80}, 40.0},
		{[]int{2, 4, 6, 8, 10, 12}, 7.0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5.0},
	}

	for _, v := range tests {
		got := CenteredAvg(v.data)
		want := v.answer
		if got != want {
			t.Error("got", got, "want", want)
		}
	}
}
