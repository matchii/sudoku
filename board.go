package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// IterationsLimit Reset execution after this number of iterations is reached
const IterationsLimit int = 1000000

type board struct {
	// Digits in cells
	data [][]int

	// Storage for temporarly used digits
	mug [][]intSlice

	// Partially filled board
	partial [][]int

	// In how many iterations board was filled
	iterations int
}

func (b *board) FillRandom() bool {
	var digits intSlice
	var prevRow, prevCol, row, col int
	for {
		row, col = b.GetNextEmptyCell()
		if row == -1 {
			b.CopyDataToPartial()
			return true // no empty cells, done
		}
		// build list of legal digits for current cell
		digits = b.GetAvailable(row, col)
		// remove from it digits previously tried
		for value := range b.mug[row][col] {
			digits.removeValue(value)
		}
		// if no digit is legal, go back to previous cell and try fill it with other (legal) digit
		if len(digits) < 1 {
			prevRow, prevCol = b.GetPreviousCell(row, col)
			b.mug[prevRow][prevCol] = append(b.mug[prevRow][prevCol], b.data[prevRow][prevCol])
			b.data[prevRow][prevCol] = 0
			// empty mug for current cell, as values remembered here were valid
			// only for previous value of preceding cell
			b.mug[row][col] = intSlice{}
			continue
		}
		b.data[row][col] = digits.randomDigit()
		b.iterations++
		if b.iterations > IterationsLimit {
			fmt.Printf(".")
			b.Reset()
		}
	}
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
		return row, col - 1
	}
	return row - 1, 8
}

func (b *board) Reset() {
	b.iterations = 0
	for r, row := range b.data {
		for c := range row {
			b.data[r][c] = 0
			b.mug[r][c] = intSlice{}
		}
	}
}

func (b *board) CopyDataToPartial() {
	for rIndex, row := range b.data {
		for cIndex := range row {
			b.partial[rIndex][cIndex] = b.data[rIndex][cIndex]
		}
	}
}

// GetAsString returns board as 81-chars long string of digits
func (b *board) GetAsString() string {
	var result [162]string
	for rIndex, row := range b.data {
		for cIndex, value := range row {
			result[9*rIndex+cIndex] = strconv.Itoa(value)
		}
	}
	for rIndex, row := range b.partial {
		for cIndex, value := range row {
			result[9*rIndex+cIndex+81] = strconv.Itoa(value)
		}
	}
	return fmt.Sprintf("%s", strings.Join(result[:], ""))
}

// WriteBoardsToFile saves given number of boards in file as 162-chars long strings
func (b *board) WriteBoardsToFile(n int, filename string) error {
	file, _ := os.Create(filename)
	defer file.Close()
	for i := 0; i < n; i++ {
		b.FillRandom()
		b.FindPartialSolution()
		file.WriteString(fmt.Sprintf("%s\n", b.GetAsString()))
		fmt.Printf("|")
		b.Reset()
	}
	return file.Sync()
}

func (b *board) FillFromString(s string) {
	digits := strings.Split(s, "")
	for rIndex, row := range b.data {
		for cIndex := range row {
			n, _ := strconv.Atoi(digits[9*rIndex+cIndex])
			b.data[rIndex][cIndex] = n
		}
	}
	for rIndex, row := range b.partial {
		for cIndex := range row {
			n, _ := strconv.Atoi(digits[9*rIndex+cIndex+81])
			b.partial[rIndex][cIndex] = n
		}
	}
}

func (b *board) FindPartialSolution() {
	for i := 0; i < 40; i++ {
		b.ClearRandomCell()
	}
}

func (b *board) ClearRandomCell() {
	r, c := b.GetRandomCell(true)
	b.partial[r][c] = 0
}

// If nonZero = true, return cell not empty on partial
func (b *board) GetRandomCell(nonZero bool) (int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var row, col int
	for {
		row = r.Intn(9)
		col = r.Intn(9)
		if nonZero && b.partial[row][col] == 0 {
			continue
		}
		break
	}
	return row, col
}

//// Printing

func (b *board) PrintFull() {
	b.Print(b.data)
}

func (b *board) PrintPartial() {
	b.Print(b.partial)
}

func (b *board) Print(grid [][]int) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	for l, line := range grid {
		if l%3 == 0 {
			tm.Println()
		}
		for c := range line {
			b.PrintCell(grid, l, c)
		}
		tm.Println()
	}
	tm.Flush()
}

func (b *board) PrintCell(grid [][]int, r, c int) {
	var cell string
	if grid[r][c] == 0 {
		cell = fmt.Sprintf(" ")
	} else {
		cell = fmt.Sprintf("%d", grid[r][c])
	}
	if c%3 == 0 {
		tm.Printf("  %s", cell)
	} else {
		tm.Printf(" %s", cell)
	}
}
