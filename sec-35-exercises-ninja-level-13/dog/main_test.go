package dog

import (
	"fmt"
	"testing"
)

// go test -v sec-35-exercises-ninja-level-13/dog/*
func TestYears(t *testing.T) {
	got := Years(5)
	want := 35

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// go test -v sec-35-exercises-ninja-level-13/dog/*
func TestYearsTwo(t *testing.T) {
	got := YearsTwo(5)
	want := 35

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(5))
	// Output:
	// 35
}

// go test -bench Years github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/dog
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

// go test -bench Years github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/dog
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
