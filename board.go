package main

import (
	tm "github.com/buger/goterm"
)

func (b *board) Fill() {
	var available = intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var current int
	for i := 0; i <= 0; i++ {
		for j := 0; j <= 8; j++ {
			current = available.randomDigit()
			b.data[i][j] = current
			available = available.removeValue(current)
		}
	}
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
