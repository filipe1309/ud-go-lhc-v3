package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("James")
	expected := "Welcome my dear James"
	if got != expected {
		t.Error("Expected", expected, "got", got)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

// go test -bench .
// go test -bench ^BenchmarkGreet$ github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/saying
// go test -bench Greet github.com/filipe1309/ud-go-lhc-v3/sec-34-testing-and-benchmarking/saying -benchmem
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
