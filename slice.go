package main

import (
	"math/rand"
	"time"
)

type intSlice []int

func (slice intSlice) removeValue(value int) intSlice {
	for {
		index := slice.indexOf(value)
		if (index < 0) {
			return slice
		}
		slice = append(slice[:index], slice[index+1:]...)
	}
}

func (slice intSlice) indexOf(value int) int {
	for i, v := range slice {
        if (v == value) {
            return i
        }
    }
    return -1
}

func (slice intSlice) randomDigit() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return slice[r.Intn(len(slice))]
}

func (slice intSlice) sum() int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

func (slice intSlice) contains(element int) bool {
	return slice.indexOf(element) > -1
}

// IntSlicesEquals checks if slices have the same unique digits, excluding 0.
// I will not make it a method, this way it can be called on board.data[index] directly.
func IntSlicesEqual(first intSlice, second intSlice) bool {
	first = first.removeValue(0)
	second = second.removeValue(0)
	if (first.sum() != second.sum()) {
		return false
	}
	for _, v := range first {
		if !second.contains(v) {
			return false
		}
	}
	for _, v := range second {
		if !first.contains(v) {
			return false
		}
	}
	return true
}
