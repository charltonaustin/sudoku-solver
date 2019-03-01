package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_checkRowsValid(t *testing.T) {
	t.Run("check valid case", func(t *testing.T) {
		m := make([]int, 81)
		for i := 0; i < 81; i++ {
			m[i] = i%9 + 1
		}
		checkRowsValid(m)
	})
}
