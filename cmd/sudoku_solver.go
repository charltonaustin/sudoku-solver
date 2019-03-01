package main

import (
	"errors"
)

func main() {
	m := make([]int, 81)
	for i := 0; i < 81; i++ {
		m[i] = i % 9 + 1
	}
	for i := 0; i < 81; i++ {
		if i%9 == 0 && i > 0 {
			println()
		}
		print(m[i])
		print(" ")
	}
}

func checkRowsValid(m []int) (bool, error) {
	if len(m) != 81 {
		return false, errors.New("invalid board")
	}
	for y := 0; y < 9; y++ {
		nums := make(map[int]bool)
		for x:=0; x < 9; x++ {
			index := (y * 9) + x
			if nums[m[index]] {
				return false, nil
			}
			nums[m[index]] = true
		}
	}
	return true, nil
}
