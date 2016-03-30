package main

import (
	"testing"
)

type testCase struct {
	slice	 intSlice
	value	 int
	expected int
}

var cases_removeValue = []testCase{
	{intSlice{}, 1, 0},
	{intSlice{5}, 5, 0},
	{intSlice{1}, 2, 1},
	{intSlice{1, 2}, 1, 1},
	{intSlice{1, 2}, 3, 2},
	{intSlice{4, 4}, 4, 0},
	{intSlice{5, 6, 6}, 6, 1},
}

func TestRemovingValue(t *testing.T) {
	for caseIndex, caseData := range cases_removeValue {
		result := caseData.slice.removeValue(caseData.value)
		if len(result) != caseData.expected {
			t.Errorf("Case #%d, expected length %s, got %s", caseIndex, caseData.expected, result)
		}
		if result.indexOf(caseData.value) > -1 {
			t.Errorf("Case #%d, value %d was not removed from slice (%s)", caseIndex, caseData.value, result)
		}
	}
}

var cases_indexOf = []testCase{
	{intSlice{}, 1, -1},
	{intSlice{1}, 1, 0},
	{intSlice{1, 2}, 1, 0},
	{intSlice{1, 2}, 2, 1},
	{intSlice{1, 2, 3}, 4, -1},
	{intSlice{1, 2, 3}, 3, 2},
}

func TestGettingIndex(t *testing.T) {
	for index, data := range cases_indexOf {
		result := data.slice.indexOf(data.value)
		if result != data.expected {
			t.Errorf("Case #%d, expected %d, got %d", index, data.expected, result)
		}
	}
}

func TestRandomDigit(t *testing.T) {
	slice := intSlice{1, 2, 7, 8}
	for i := 0; i < 100; i++ {
		result := slice.randomDigit()
		if slice.indexOf(result) < 0 {
			t.Errorf("Value from outside of slice was generated: %d", result)
		}
	}
}
