package main

import (
	"fmt"
	"os"
	"strconv"
)

// Usage:
//
//		sudoku gen <n> <filename>
//
// Generates n boards and saves them to file filename
//
//		sudoku read <filename>
//
// Reads first board from file <filename> and displays it as a grid
//
// If called without parameters, print one random board and exit.
func main() {
	if len(os.Args) == 1 {
		b := newBoard()
		b.FillRandom()
		b.FindPartialSolution()
		b.PrintFull()
		b.PrintPartial()
		os.Exit(0)
	}
	if os.Args[1] == "gen" {
		generateToFile()
		os.Exit(0)
	}
	if os.Args[1] == "read" {
		printFromFile()
		os.Exit(0)
	}
}

func generateToFile() {
	b := newBoard()
	number := 1
	filename := "sudoku.txt"
	if len(os.Args) > 2 {
		number, _ = strconv.Atoi(os.Args[2])
	}
	if len(os.Args) > 3 {
		filename = os.Args[3]
	}
	if writeError := b.WriteBoardsToFile(number, filename); writeError != nil {
		fmt.Printf("\nError when writing to file: %s\n", writeError)
	} else {
		fmt.Printf("\n%d board(s) generated and stored in file %s\n", number, filename)
	}
}

func printFromFile() {
	filename := "sudoku.txt"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	data := make([]byte, 162)
	f, _ := os.Open(filename)
	count, _ := f.Read(data)
	b := newBoard()
	b.FillFromString(string(data[:count]))
	b.PrintPartial()
	b.ResolvePartial()
	b.PrintFull()
	b.PrintPartial()
}

func newBoard() board {
	data := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		data[i] = make([]int, 9)
	}
	mug := make([][]intSlice, 9)
	for i := 0; i <= 8; i++ {
		mug[i] = make([]intSlice, 9)
	}
	partial := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		partial[i] = make([]int, 9)
	}
	fixed := make([][]bool, 9)
	for i := 0; i <= 8; i++ {
		fixed[i] = make([]bool, 9)
	}
	b := board{data, mug, partial, fixed, 0}
	return b
}
