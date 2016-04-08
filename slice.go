package main

import (
	"math/rand"
	"time"
)

type intSlice []int

func (s *intSlice) removeValue(value int) {
	for {
		index := s.indexOf(value)
		if (index < 0) {
			return
		}
		tmp := *s
		*s = append(tmp[:index], tmp[index+1:]...)
	}
}

func (s *intSlice) indexOf(value int) int {
	for i, v := range *s {
        if (v == value) {
            return i
        }
    }
    return -1
}

func (s *intSlice) randomDigit() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return (*s)[r.Intn(len(*s))]
}

func (s *intSlice) sum() int {
	var sum int
	for _, v := range *s {
		sum += v
	}
	return sum
}

func (s *intSlice) contains(element int) bool {
	return s.indexOf(element) > -1
}

// IntSlicesEquals checks if slices have the same unique digits, excluding 0.
// I will not make it a method, this way it can be called on board.data[index] directly.
// Parameters are passed by value on purpose.
func IntSlicesEqual(first intSlice, second intSlice) bool {
	first.removeValue(0)
	second.removeValue(0)
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
