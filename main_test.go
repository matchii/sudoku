package main

import (
	"testing"
)

type TestCase struct {
	needle   int
	haystack []int
	expected bool
}

var testCases = []TestCase{
	{1, []int{}, false},
	{1, []int{1}, true},
	{1, []int{2}, false},
	{1, []int{1, 2, 3}, true},
	{1, []int{7, 8, 9}, false},
	{1, []int{1, 1, 2}, true},
	{1, []int{5, 5, 5}, false},
}

func TestInSlice(t *testing.T) {
	for i, c := range testCases {
		var result = InSlice(c.needle, c.haystack)
		if result != c.expected {
			t.Errorf("For case %d expected %s, got %s", i, c.expected, result)
		}
	}
}
