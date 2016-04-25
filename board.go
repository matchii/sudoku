package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	tm "github.com/buger/goterm"
)

// Reset execution after this number of iterations is reached
const ITERATIONS_LIMIT int = 1000000

type board struct {
	// Digits in cells
	data [][]int

	// In how many iterations board was filled
	iterations int
}

func (b *board) Fill() bool {
	var digits intSlice
	var prevRow, prevCol, row, col int
	mug := BuildMug()
	for {
		row, col = b.GetNextEmptyCell()
		if row == -1 {
			return true // no empty cells, done
		}
		// build list of legal digits for current cell
		digits = b.GetAvailable(row, col)
		// remove from it digits previously tried
		for value := range mug[row][col] {
			digits.removeValue(value)
		}
		// if no digit is legal, go back to previous cell and try fill it with other (legal) digit
		if len(digits) < 1 {
			prevRow, prevCol = b.GetPreviousCell(row, col)
			mug[prevRow][prevCol] = append(mug[prevRow][prevCol], b.data[prevRow][prevCol])
			b.data[prevRow][prevCol] = 0
			// empty mug for current cell, as values remembered here were valid
			// only for previous value of preceding cell
			mug[row][col] = intSlice{}
			continue
		}
		b.data[row][col] = digits.randomDigit()
		b.iterations++
		if b.iterations > ITERATIONS_LIMIT {
			b.Empty()
		}
	}
}

func BuildMug() [][]intSlice {
	mug := make([][]intSlice, 9)
	for i := 0; i <= 8; i++ {
		mug[i] = make([]intSlice, 9)
	}
	return mug
}

// GetAvailable returns slice of digits that can be inserted at given position.
func (b *board) GetAvailable(row int, col int) intSlice {
	tmp := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// row
	for x := 0; x <= 8; x++ {
		tmp.removeValue(b.data[row][x])
	}
	// column
	for y := 0; y <= 8; y++ {
		tmp.removeValue(b.data[y][col])
	}
	// block
	blockFirstRow := row / 3 * 3
	blockFirstCol := col / 3 * 3
	for i := blockFirstRow; i <= blockFirstRow+2; i++ {
		for j := blockFirstCol; j <= blockFirstCol+2; j++ {
			tmp.removeValue(b.data[i][j])
		}
	}
	return tmp
}

// GetNextEmptyCell returns index (row, col) of the first empty cell,
// or (-1, -1) if all cells are filled.
func (b *board) GetNextEmptyCell() (int, int) {
	for rowIdx, row := range b.data {
		for colIdx := range row {
			if b.data[rowIdx][colIdx] == 0 {
				return rowIdx, colIdx
			}
		}
	}
	return -1, -1
}

func (b *board) GetPreviousCell(row, col int) (int, int) {
	if row == 0 && col == 0 {
		return -1, -1
	}
	if col > 0 {
		return row, col-1
	}
	return row-1, 8
}

func (b *board) Empty() {
	for r, row := range b.data {
		for c := range row {
			b.data[r][c] = 0
		}
	}
}

func (b *board) Print() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	for l, line := range b.data {
		if l % 3 == 0 {
			tm.Println()
		}
		for c, _ := range line {
			if c % 3 == 0 {
				tm.Printf("  %d", b.data[l][c])
			} else {
				tm.Printf(" %d", b.data[l][c])
			}
		}
		tm.Println()
	}
	tm.Println()
	tm.Printf("Board generated in %d iterations", b.iterations)
	tm.Flush()
}

// GetAsString returns board as 81-chars long string of digits
func (b *board) GetAsString() string {
	var result [81]string
	for rIndex, row := range b.data {
		for cIndex, value := range row {
			result[9 * rIndex + cIndex] = strconv.Itoa(value)
		}
	}
	return fmt.Sprintf("%s", strings.Join(result[:], ""))
}

// WriteBoardsToFile saves given number of boards in file as 81-chars long strings
func (b *board) WriteBoardsToFile(n int, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	for i := 0; i < n; i++ {
		b.Fill()
		file.WriteString(fmt.Sprintf("%s\n", b.GetAsString()))
		b.Empty()
	}
	file.Sync()
}
