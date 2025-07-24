package main

import "slices"

func CheckNoContains(slice1, slice2 []string) []string {
	var newSlice []string
	for _, value := range slice1 {
		if !slices.Contains(slice2, value) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
