package main

import (
	"fmt"
	"os"
	"strconv"
)

// Parameters are:
//	- number of boards to generate
//	- file to store them, default "sudoku.txt"
//
// If called without parameters, print one board and exit.
func main() {
	b := NewBoard()
	var number int
	var filename string
	if len(os.Args) > 1 {
		number, _ = strconv.Atoi(os.Args[1])
	} else {
		b.Fill()
		b.Print()
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		filename = os.Args[2]
	} else {
		filename = "sudoku.txt"
	}
	b.WriteBoardsToFile(number, filename)
	fmt.Printf("%d board(s) generated and stored in file %s\n", number, filename)
}

func NewBoard() board {
	data := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		data[i]  = make([]int, 9)
	}
	b := board{data, 0}
	return b
}
