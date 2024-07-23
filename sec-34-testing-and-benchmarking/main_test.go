package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	result := mySum(2, 3)
	expected := 5
	if result != expected {
		t.Error("Expected", expected, "got", result)
	}
}
