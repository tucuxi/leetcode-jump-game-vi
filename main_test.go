package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxResult(t *testing.T) {
	t.Run("nums=[1,-1,-2,4,-7,3] k=2", func(t *testing.T) {
		res := maxResult([]int{1, -1, -2, 4, -7, 3}, 2)
		assert.Equal(t, 7, res)
	})
}
