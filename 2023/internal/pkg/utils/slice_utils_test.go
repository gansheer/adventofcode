package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestIntersection(t *testing.T) {
	left := []int{83, 86, 6, 31, 17, 9, 48, 53}
	right := []int{41, 48, 83, 86, 17}

	intersect := Intersection(left, right)

	assert.Equal(t, 4, len(intersect))
}

func TestMinIntSlice(t *testing.T) {
	slice := []int{83, 86, 6, 31, 17, 9, 48, 53}
	min := MinIntSlice(slice)
	assert.Equal(t, 6, min)

	minEmpty := MinIntSlice([]int{})
	assert.Equal(t, 0, minEmpty)
}
