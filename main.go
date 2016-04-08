package main

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
	b := board{data, 0}
	return b
}
