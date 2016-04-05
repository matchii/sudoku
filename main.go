package main

type board struct {
	data [][]int
}

func main() {
	b := NewBoard()
	b.Fill()
	b.Print()
}

func NewBoard() board {
	data := make([][]int, 9)
	for i := 0; i <= 8; i++ {
		data[i]  = make([]int, 9)
	}
	b := board{data}
	return b
}
