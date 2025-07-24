package main

import (
	"slices"
	"testing"
)

func Test_CheckContains(t *testing.T) {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}
	slice3 := []int{1, 2, 3, 4}
	slice4 := []int{5, 6, 7, 8}

	resultCheck, resultSlice := CheckContains(slice1, slice2)
	if resultCheck != true {
		t.Error("invalid check param must be true")
	}
	if len(resultSlice) != 2 || !slices.Contains(resultSlice, 64) || !slices.Contains(resultSlice, 3) {
		t.Errorf("slices is not equal returned: %v", resultSlice)
	}

	resultCheck, resultSlice = CheckContains(slice3, slice4)
	if resultCheck != false {
		t.Error("invalid check param must be false")
	}
	if resultSlice != nil {
		t.Errorf("slices is not equal not nil returned: %v", resultSlice)
	}
}
