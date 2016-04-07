package main

import (
	"github.com/stvp/assert"
	"testing"
)

func TestGetAvailable(t *testing.T) {
	b := NewBoard()
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		"C1: All should be available")
	b.data[0][0] = 1
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{2, 3, 4, 5, 6, 7, 8, 9}),
		"C2: All but 1 should be available")
	b.data[0][1] = 2
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{3, 4, 5, 6, 7, 8, 9}),
		"C3: All but 1, 2 should be available")
	b.data[1][0] = 3
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{4, 5, 6, 7, 8, 9}),
		"C4: All but 1, 2, 3 should be available")
	b.data[1][1] = 4
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{5, 6, 7, 8, 9}),
		"C5: All but 1, 2, 3, 4 should be available")
	b.data[3][3] = 5
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{5, 6, 7, 8, 9}),
		"C6: All but 1, 2, 3, 4 should be available")
	b.data[8][0] = 6
	b.data[4][0] = 6
	b.data[0][4] = 7
	b.data[0][8] = 7
	assert.True(t,
		IntSlicesEqual(b.GetAvailable(0, 0), intSlice{5, 8, 9}),
		"C7: All but 1, 2, 3, 4, 5, 6, 7 should be available")
}
