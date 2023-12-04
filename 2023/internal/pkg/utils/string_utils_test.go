package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindFirst(t *testing.T) {
	input := "xtwone3four"
	token, index := FindFirst(input, []string{"one", "two", "four", "five"})

	assert.Equal(t, "two", token)
	assert.Equal(t, 1, index)
}
func TestReplaceStringNumberByDigit(t *testing.T) {
	input := "xtwone3four"
	output := ReplaceStringNumberByDigit(input)

	assert.Equal(t, "xtw2o1ne3fo4ur", output)
}
