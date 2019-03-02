package board

import (
	"errors"
	"strconv"
	"strings"
)

type Board struct {
	board []int
}

func NewBoard() Board {
	m := make([]int, 81)
	return Board{m}
}

func (b *Board) Update(x, y, value int) error {
	if x < 0 || x >= 9 {
		return errors.New("invalid x")
	}
	if y < 0 || y >= 9 {
		return errors.New("invalid y")
	}
	if value < 1 || value >= 10 {
		return errors.New("invalid value")
	}

	index := (y * 9) + x
	b.board[index] = value
	return nil
}

func (b *Board) Get(x, y int) (int, error) {
	index := (y * 9) + x
	if index < 0 || index >= 81 {
		return 0, errors.New("invalid coordinates")
	}
	return b.board[index], nil
}

func (b Board) checkRowsValid() bool {
	for y := 0; y < 9; y++ {
		nums := make(map[int]bool)
		for x := 0; x < 9; x++ {
			index := (y * 9) + x
			if nums[b.board[index]] {
				return false
			}
			nums[b.board[index]] = true
		}
	}
	return true
}

func (b Board) checkColumnsValid() bool {
	for x := 0; x < 9; x++ {
		nums := make(map[int]bool)
		for y := 0; y < 9; y++ {
			index := (y * 9) + x
			if nums[b.board[index]] {
				return false
			}
			nums[b.board[index]] = true
		}
	}
	return true
}

func (b Board) checkValidSections() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nums := make(map[int]bool)
			for y := i * 3; y < i*3+3; y++ {
				for x := j * 3; x < j*3+3; x++ {
					index := (y * 9) + x
					if nums[b.board[index]] {
						return false
					}
					nums[b.board[index]] = true
				}
			}
		}
	}
	return true
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < 81; i++ {
		if i%9 == 0 && i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(strconv.Itoa(b.board[i]))
		sb.WriteString(" ")
	}
	return sb.String()
}
