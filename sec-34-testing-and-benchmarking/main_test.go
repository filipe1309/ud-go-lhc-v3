package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	got := mySum(2, 3)
	expected := 5
	if got != expected {
		t.Error("Expected", expected, "got", got)
	}
}

func TestMySumWithTableTestsClass246(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		// test{[]int{21, 21}, 42},
		{[]int{21, 21}, 42},
		{[]int{2, 3}, 5},
		{[]int{4, 7}, 11},
		{[]int{5, 9, 1}, 15},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		result := mySum(v.data...)
		if result != v.answer {
			t.Error("Expected", v.answer, "got", result)
		}
	}
}
