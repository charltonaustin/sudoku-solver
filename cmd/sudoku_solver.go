package main

import (
	"fmt"
	"math/rand"
	"sudoku_solver/board"
	"time"
)

func main() {
	b := board.NewBoard()
	rand.Seed(time.Now().UTC().UnixNano())
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for x := 0; x < 9; x++ {
		index := rand.Intn(len(numbers))
		newValue := numbers[index]
		b.Update(x, 0, newValue)
		numbers = append(numbers[:index], numbers[index+1:]...)
	}
	backTrace(b)
	fmt.Println(b)
}

func backTrace(b *board.Board) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			value, err := b.Get(x, y)
			if err != nil {
				panic(err)
			}
			if value == 0 {
				for number := 1; number < 10; number++ {
					err := b.Update(x, y, number)
					if err != nil {
						panic(err)
					}
					if b.IsValid() {
						if backTrace(b) {
							return true
						}
						b.Update(x, y, 0)
					} else {
						b.Update(x, y, 0)
					}

				}
				return false
			}
		}
	}
	return true
}