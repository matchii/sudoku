package main

import (
	tm "github.com/buger/goterm"
)

type board struct {
	data [][]int
}

func main() {
	b := NewBoard()
	b.Print()
}

func NewBoard() board {
	data := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		data[i]  = make([]int, 9)
	}
	b := board{data}
	b.Fill()
	return b
}

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
