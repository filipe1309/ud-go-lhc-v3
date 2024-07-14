package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(5, 5)
	want := 10
	if got != want {
		t.Errorf("Sum was incorrect, got %d, want: %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(42, 16)
	want := 26
	if got != want {
		log.Fatalf("error - want: %d, got: %d", want, got)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(42, 16, add)
	want := 58
	if got != want {
		log.Fatalf("error - want: %d, got: %d", want, got)
	}
	got = doMath(42, 16, subtract)
	want = 26
	if got != want {
		log.Fatalf("error - want: %d, got: %d", want, got)
	}
}

func TestSquareArea(t *testing.T) {
	s := square{
		length: 2,
		width: 3,
	}
	got := s.area()
	want := 6.0
	if got != want {
		log.Fatalf("error - want: %v, got: %v", want, got)
	}
}

func TestParadise(t *testing.T) {
	got := paradise("Buenos Aires")
	want := "My idea of paradise is Buenos Aires"
	if got != want {
		log.Fatalf("error - want: %v, got: %v", want, got)
	}
}
