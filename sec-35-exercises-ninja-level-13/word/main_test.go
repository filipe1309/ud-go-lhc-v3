package word

import (
	"fmt"
	"testing"

	"github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/quote"
)

// go test -bench Count github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/word
func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

// go test -bench Count github.com/filipe1309/ud-go-lhc-v3/sec-35-exercises-ninja-level-13/word
func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

// go test -v sec-35-exercises-ninja-level-13/word/*
func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))
	// Output:
	// 1325
}

// go test -v sec-35-exercises-ninja-level-13/word/*
func ExampleUseCount() {
	fmt.Println(UseCount("one two two three")["two"])
	// Output:
	// 2
}

// go test -v sec-35-exercises-ninja-level-13/word/*
func TestCount(t *testing.T) {
	got := Count("one two three")
	want := 3
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// go test -v sec-35-exercises-ninja-level-13/word/*
func TestUseCount(t *testing.T) {
	got := UseCount("one two three three three")
	want := map[string]int{
		"one":   1,
		"two":   1,
		"three": 3,
	}
	for k, v := range got {
		if v != want[k] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
