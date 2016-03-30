package main

import (
	//"fmt"
	tm "github.com/buger/goterm"
	"math/rand"
	"time"
)

type board struct {
	data [][]int
	mask [][]bool
	full_nine []int
}

func main() {
	b := NewBoard()
	b.Print()
}

func NewBoard() board {
	full_nine := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	data := make([][]int, 9)
	mask := make([][]bool, 9)
	for i := 0; i <= 8; i++ {
		data[i]  = make([]int, 9)
		mask[i]  = make([]bool, 9)
	}
	b := board{data, mask, full_nine}
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

func RandomDigit(exclude []int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		try := r.Intn(9)
		if InSlice(try, exclude) {
			continue
		}
		return try;
	}
}

func InSlice(needle int, haystack []int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
