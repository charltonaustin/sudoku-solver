package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"sudoku_solver/board"
	"time"
)

func main() {
	usage := "path to a file with a sudoku puzzle, if none is provided then this will generate a random valid solved board"
	filePath := flag.String("file", "", usage)
	flag.Parse()
	if *filePath == "" {
		rand.Seed(time.Now().UTC().UnixNano())
		b := board.NewBoard()
		backTrace(b)
		fmt.Println(fmt.Sprintf("is valid: %v", b.IsValid()))
		fmt.Println(b)
		return
	}
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}
	split := strings.Split(strings.Trim(string(data), "\n"), "\n")

	b := board.NewBoard()
	for y, column := range split {
		rows := strings.Split(column, " ")
		for x, number := range rows {
			value, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			b.Update(x, y, value)
		}
	}
	valid := b.IsValid()
	if !valid {
		panic("puzzle is invalid")
	}
	fmt.Println("soving puzzle")
	fmt.Println(b)
	backTrace(b)
	fmt.Println("solved puzzle")
	fmt.Println(b)

}

func randomNumbers() []int {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
	return numbers
}

func backTrace(b *board.Board) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			value, err := b.Get(x, y)
			if err != nil {
				panic(err)
			}
			if value == 0 {
				for _, number := range randomNumbers() {
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