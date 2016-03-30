package main

import (
	"math/rand"
	"time"
)

type intSlice []int

func (slice intSlice) removeValue(value int) intSlice {
	index := slice.indexOf(value)
	if (index < 0) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
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
