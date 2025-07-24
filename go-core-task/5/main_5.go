package main

import (
	"slices"
)

func CheckContains(slice1, slice2 []int) (bool, []int) {
	var newSlice []int
	for _, value := range slice1 {
		if slices.Contains(slice2, value) {
			newSlice = append(newSlice, value)
		}
	}

	if len(newSlice) != 0 {
		return true, newSlice
	}

	return false, nil
}
