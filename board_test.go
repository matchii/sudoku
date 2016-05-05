package main

import (
	"github.com/stvp/assert"
	"testing"
)

func TestGetAvailable(t *testing.T) {
	b := newBoard()
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

func TestGetNextEmptyCell(t *testing.T) {
	b := newBoard()
	var row, col int
	// empty board, expect 0, 0
	row, col = b.GetNextEmptyCell()
	assert.Equal(t, 0, row)
	assert.Equal(t, 0, col)
	// get next from first row
	b.data[0][0] = 1
	row, col = b.GetNextEmptyCell()
	assert.Equal(t, 0, row)
	assert.Equal(t, 1, col)
	// ignore this, it is in second row and first still has empty cells
	b.data[1][0] = 1
	row, col = b.GetNextEmptyCell()
	assert.Equal(t, 0, row)
	assert.Equal(t, 1, col)
	// fill entire first row
	for i := 1; i <= 8; i++ {
		b.data[0][i] = i
	}
	// should return first empty from second row (1, 1)
	row, col = b.GetNextEmptyCell()
	assert.Equal(t, 1, row)
	assert.Equal(t, 1, col)
	// fill everything
	for r := 1; r <= 8; r++ {
		for c := 0; c <= 8; c++ {
			b.data[r][c] = 1
		}
	}
	// board filled, expect -1, -1
	row, col = b.GetNextEmptyCell()
	assert.Equal(t, -1, row)
	assert.Equal(t, -1, col)
}

func TestGetPreviousCell(t *testing.T) {
	b := newBoard()
	var row, col int

	// First cell, there is no previous one
	row, col = b.GetPreviousCell(0, 0)
	assert.Equal(t, -1, row)
	assert.Equal(t, -1, col)
	// Previous in the same row
	row, col = b.GetPreviousCell(0, 5)
	assert.Equal(t, 0, row)
	assert.Equal(t, 4, col)
	// Previous non-fixed in the same row
	b.fixed[0][4] = true
	row, col = b.GetPreviousCell(0, 5)
	assert.Equal(t, 0, row)
	assert.Equal(t, 3, col)
	// Last from previous row
	row, col = b.GetPreviousCell(8, 0)
	assert.Equal(t, 7, row)
	assert.Equal(t, 8, col)
	// Last non-fixed from previous row
	b.fixed[8][0] = true
	row, col = b.GetPreviousCell(8, 1)
	assert.Equal(t, 7, row)
	assert.Equal(t, 8, col)
}

func TestEmpty(t *testing.T) {
	b := newBoard()
	b.data[0][0] = 1
	b.data[2][2] = 1
	b.data[5][5] = 1
	b.data[7][7] = 1
	b.Reset()
	assert.Equal(t, 0, b.data[0][0])
	assert.Equal(t, 0, b.data[2][2])
	assert.Equal(t, 0, b.data[5][5])
	assert.Equal(t, 0, b.data[7][7])
}
