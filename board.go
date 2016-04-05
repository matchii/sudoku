package main

import (
	"fmt"
	tm "github.com/buger/goterm"
)

func (b *board) Fill() {
	for i := 0; i <= 0; i++ {
		for j := 0; j <= 8; j++ {
			b.data[i][j] = b.GetAvailable(i, j).randomDigit()
		}
	}
}

func (b *board) GetAvailable(row int, col int) intSlice {
	tmp := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for x := 0; x <= 8; x++ {
		tmp = tmp.removeValue(b.data[row][x])
	}
	for y := 0; y <= 8; y++ {
		tmp = tmp.removeValue(b.data[y][col])
	}
	if (len(tmp) == 0) {
		panic(fmt.Sprintf("No digit available at x=%d y=%d", col, row))
	}
	return tmp
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
