package main

import (
	"testing"
)

func TestGetAvailable(t *testing.T) {
	b := NewBoard()
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		"C1: All should be available")
	b.data[0][0] = 1
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{2, 3, 4, 5, 6, 7, 8, 9}),
		"C2: All but 1 should be available")
	b.data[0][1] = 2
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{3, 4, 5, 6, 7, 8, 9}),
		"C3: All but 1, 2 should be available")
	b.data[1][0] = 3
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{4, 5, 6, 7, 8, 9}),
		"C4: All but 1, 2, 3 should be available")
	b.data[1][1] = 4
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{4, 5, 6, 7, 8, 9}),
		"C5: All but 1, 2, 3 should be available")
	b.data[8][0] = 5
	b.data[4][0] = 5
	b.data[0][4] = 6
	b.data[0][8] = 6
	assert(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{4, 7, 8, 9}),
		"C5: All but 1, 2, 3, 5, 6 should be available")
}

func assert(t *testing.T, param bool, message string) {
	if (!param) {
		t.Errorf("Assertion failed: %s", message)
	}
}
