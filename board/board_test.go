package board

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkRowsValid(t *testing.T) {
	t.Run("check valid case", func(t *testing.T) {
		m := make([]int, 81)
		for i := 0; i < 81; i++ {
			m[i] = i%9 + 1
		}
		board := Board{board: m}
		b := board.checkRowsValid()
		assert.True(t, b)
	})

	t.Run("check invalid case", func(t *testing.T) {
		m := make([]int, 81)
		for i := 0; i < 81; i++ {
			m[i] = i%9 + 1
		}
		m[80] = 1
		board := Board{board: m}
		b := board.checkRowsValid()
		assert.False(t, b)
	})
}

func Test_checkColumnsValid(t *testing.T) {
	t.Run("check valid case", func(t *testing.T) {
		m := make([]int, 81)
		num := 0
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				index := (y * 9) + x
				m[index] = num
			}
			num++
		}
		board := Board{board: m}
		b := board.checkColumnsValid()
		assert.True(t, b)
	})

	t.Run("check invalid case", func(t *testing.T) {
		m := make([]int, 81)
		num := 0
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				index := (y * 9) + x
				m[index] = num
			}
			num++
		}
		m[80] = 1
		board := Board{board: m}
		b := board.checkColumnsValid()
		assert.False(t, b)
	})
}

func Test_checkValidSections(t *testing.T) {
	t.Run("check valid case", func(t *testing.T) {
		m := make([]int, 81)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				num := 0
				for y := i * 3; y < i*3+3; y++ {
					for x := j * 3; x < j*3+3; x++ {
						index := (y * 9) + x
						m[index] = num
						num++
					}
				}
			}
		}
		board := Board{board: m}
		fmt.Println(board.String())
		b := board.checkValidSections()
		assert.True(t, b)
	})

	t.Run("check invalid case", func(t *testing.T) {
		m := make([]int, 81)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				num := 0
				for y := i * 3; y < i*3+3; y++ {
					for x := j * 3; x < j*3+3; x++ {
						index := (y * 9) + x
						m[index] = num
						num++
					}
				}
			}
		}
		m[80] = 1
		board := Board{board: m}
		b := board.checkValidSections()
		assert.False(t, b)
	})
}
