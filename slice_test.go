package main

import (
	"github.com/stvp/assert"
	"testing"
)

type TestCaseInt struct {
	slice    intSlice
	value    int
	expected int
}

type TestCaseBool struct {
	slice    intSlice
	value    int
	expected bool
}

var casesRemoveValue = []TestCaseInt{
	{intSlice{}, 1, 0},
	{intSlice{5}, 5, 0},
	{intSlice{1}, 2, 1},
	{intSlice{1, 2}, 1, 1},
	{intSlice{1, 2}, 3, 2},
	{intSlice{4, 4}, 4, 0},
	{intSlice{5, 6, 6}, 6, 1},
}

func TestRemovingValue(t *testing.T) {
	for _, data := range casesRemoveValue {
		data.slice.removeValue(data.value)
		assert.False(t, len(data.slice) != data.expected)
		assert.False(t, data.slice.indexOf(data.value) > -1)
	}
}

var casesIndexOf = []TestCaseInt{
	{intSlice{}, 1, -1},
	{intSlice{1}, 1, 0},
	{intSlice{1, 2}, 1, 0},
	{intSlice{1, 2}, 2, 1},
	{intSlice{1, 2, 3}, 4, -1},
	{intSlice{1, 2, 3}, 3, 2},
}

func TestGettingIndex(t *testing.T) {
	for _, data := range casesIndexOf {
		assert.Equal(t, data.expected, data.slice.indexOf(data.value))
	}
}

var casesSum = []TestCaseInt{
	{intSlice{}, 0, 0},
	{intSlice{0}, 0, 0},
	{intSlice{1}, 0, 1},
	{intSlice{1, 2, 3}, 0, 6},
	{intSlice{4, 4, 4}, 0, 12},
}

func TestSum(t *testing.T) {
	for _, data := range casesSum {
		assert.Equal(t, data.expected, data.slice.sum())
	}
}

var casesContains = []TestCaseBool{
	{intSlice{}, 1, false},
	{intSlice{1}, 1, true},
	{intSlice{1}, 2, false},
	{intSlice{1, 1}, 1, true},
	{intSlice{1, 1}, 2, false},
	{intSlice{1, 2, 3}, 2, true},
}

func TestContains(t *testing.T) {
	for _, data := range casesContains {
		assert.Equal(t, data.expected, data.slice.contains(data.value))
	}
}

func TestRandomDigit(t *testing.T) {
	slice := intSlice{1, 2, 7, 8}
	for i := 0; i < 100; i++ {
		assert.False(t, slice.indexOf(slice.randomDigit()) < 0)
	}
}
