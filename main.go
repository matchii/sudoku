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
		b := NewBoard()
		b.FillRandom()
		b.PrintFull()
		os.Exit(0)
	}
	if os.Args[1] == "gen" {
		GenerateToFile()
		os.Exit(0)
	}
	if os.Args[1] == "read" {
		PrintFromFile()
		os.Exit(0)
	}
}

func GenerateToFile() {
	b := NewBoard()
	number := 1
	filename := "sudoku.txt"
	if len(os.Args) > 2 {
		number, _ = strconv.Atoi(os.Args[2])
	}
	if len(os.Args) > 3 {
		filename = os.Args[3]
	}
	b.WriteBoardsToFile(number, filename)
	fmt.Printf("\n%d board(s) generated and stored in file %s\n", number, filename)
}

func PrintFromFile() {
	filename := "sudoku.txt"
	if len(os.Args) > 2 {
		filename = os.Args[2]
	}
	data := make([]byte, 81)
	f, _ := os.Open(filename)
	count, _ := f.Read(data)
	b := NewBoard()
	b.FillFromString(string(data[:count]))
	b.PrintFull()
}

func NewBoard() board {
	data := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		data[i]  = make([]int, 9)
	}
	mug := make([][]intSlice, 9)
	for i := 0; i <= 8; i++ {
		mug[i] = make([]intSlice, 9)
	}
	partial := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		partial[i] = make([]int, 9)
	}
	b := board{data, mug, partial, 0}
	return b
}
