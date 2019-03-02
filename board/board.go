package board

import (
	"errors"
	"strconv"
	"strings"
)

type Board struct {
	board []int
}

func NewBoard() *Board {
	m := make([]int, 81)
	return &Board{m}
}

func (b *Board) Update(x, y, value int) error {
	if x < 0 || x >= 9 {
		return errors.New("invalid x")
	}
	if y < 0 || y >= 9 {
		return errors.New("invalid y")
	}
	if value < 0 || value >= 10 {
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
		if !b.checkRowValid(y) {
			return false
		}
	}
	return true
}

func (b Board) checkRowValid(row int) bool {
	nums := make(map[int]bool)
	for x := 0; x < 9; x++ {
		index := (row * 9) + x
		if nums[b.board[index]] && b.board[index] != 0 {
			return false
		}
		nums[b.board[index]] = true
	}
	return true
}

func (b Board) checkColumnsValid() bool {
	for x := 0; x < 9; x++ {
		if !b.checkColumnValid(x) {
			return false
		}
	}
	return true
}

func (b Board) checkColumnValid(column int) bool {
	nums := make(map[int]bool)
	for y := 0; y < 9; y++ {
		index := (y * 9) + column
		if nums[b.board[index]] && b.board[index] != 0 {
			return false
		}
		nums[b.board[index]] = true
	}
	return true
}

func (b Board) checkValidSections() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !b.checkValidSection(i, j) {
				return false
			}
		}
	}
	return true
}

func (b Board) checkValidSection(i, j int) bool {
	nums := make(map[int]bool)
	for y := i * 3; y < i*3+3; y++ {
		for x := j * 3; x < j*3+3; x++ {
			index := (y * 9) + x
			if nums[b.board[index]] && b.board[index] != 0 {
				return false
			}
			nums[b.board[index]] = true
		}
	}
	return true
}

func(b *Board) IsValid()bool {
	return b.checkRowsValid() && b.checkColumnsValid() && b.checkValidSections()
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < 81; i++ {
		if i%9 == 0 && i > 0 {
			sb.WriteString("\n")
		}
		value := strconv.Itoa(b.board[i])
		if value == "0" {
			value = "_"
		}
		sb.WriteString(value)
		sb.WriteString(" ")
	}
	return sb.String()
}