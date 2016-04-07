package main

import (
	tm "github.com/buger/goterm"
)

func (b *board) Fill() {
	for i := 0; i <= 0; i++ {
		for j := 0; j <= 8; j++ {
			b.data[i][j] = b.GetAvailable(i, j).randomDigit()
		}
	}
}

// GetAvailable returns slice of digits that can be inserted at given position.
func (b *board) GetAvailable(row int, col int) intSlice {
	tmp := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// row
	for x := 0; x <= 8; x++ {
		tmp = tmp.removeValue(b.data[row][x])
	}
	// column
	for y := 0; y <= 8; y++ {
		tmp = tmp.removeValue(b.data[y][col])
	}
	// block
	blockFirstRow := row / 3 * 3
	blockFirstCol := col / 3 * 3
	for i := blockFirstRow; i <= blockFirstRow+2; i++ {
		for j := blockFirstCol; j <= blockFirstCol+2; j++ {
			tmp = tmp.removeValue(b.data[i][j])
		}
	}
	return tmp
}

// GetNextEmpty returns index (row, col) of the first empty cell,
// or (-1, -1) if all cells are filled.
func (b *board) GetNextEmpty() (int, int) {
	for rowIdx, row := range b.data {
		for colIdx := range row {
			if b.data[rowIdx][colIdx] == 0 {
				return rowIdx, colIdx
			}
		}
	}
	return -1, -1
}

func (b *board) Print() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	for l, line := range b.data {
		for c, _ := range line {
			tm.Printf(" %d", b.data[l][c])
		}
		tm.Println()
	}
	tm.Flush()
}
