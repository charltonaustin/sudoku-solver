package main

import (
	"fmt"
	board2 "sudoku_solver/board"
)

func main() {
	board := board2.NewBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := 0
			for y := i * 3; y < i*3+3; y++ {
				for x := j * 3; x < j*3+3; x++ {
					board.Update(x, y, num)
					num++
				}
			}
		}
	}
	fmt.Println(board.String())
}

