package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSample(t *testing.T) {
	result, err := computeSimple("testdata/sample1.txt")

	assert.Nil(t, err)
	assert.Equal(t, "4361", result)

}

func TestSimpleData(t *testing.T) {
	result, err := computeSimple("testdata/data.txt")

	assert.Nil(t, err)
	assert.Equal(t, "533784", result)

}

func TestComplexSample(t *testing.T) {
	result, err := computeComplex("testdata/sample2.txt")

	assert.Nil(t, err)
	assert.Equal(t, "467835", result)

}

func TestComplexData(t *testing.T) {
	result, err := computeComplex("testdata/data.txt")

	assert.Nil(t, err)
	assert.Equal(t, "0", result)

}
