package main

import (
	"slices"
	"testing"
)

func Test_CheckNoContains(t *testing.T) {
	slice1 := []string{"asd", "zxc", "qwe", "dfg"}
	slice2 := []string{"asd", "cvb", "uio"}
	correctResultSlice := []string{"zxc", "qwe", "dfg"}

	resultSlice := CheckNoContains(slice1, slice2)
	if !slices.Equal(resultSlice, correctResultSlice) {
		t.Errorf("CheckNoContains(%v, %v) = %v; got %v", slice1, slice2, resultSlice, correctResultSlice)
	}
}
