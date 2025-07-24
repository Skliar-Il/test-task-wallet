package main

import (
	"reflect"
	"testing"
)

func Test_EvenSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := evenSlice(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("evenSlice(%v) = %v; want %v", input, result, expected)
	}
}

func Test_AddElements(t *testing.T) {
	input := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(input, element)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("addElements(%v, %d) = %v; want %v", input, element, result, expected)
	}
}

func Test_CopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	result := copySlice(input)

	if !reflect.DeepEqual(result, input) {
		t.Errorf("copySlice(%v) = %v; want %v", input, result, input)
	}

	input[0] = 99
	if input[0] == result[0] {
		t.Errorf("copySlice returned slice affected by changes in original slice")
	}
}

func Test_RemoveElement(t *testing.T) {
	input := []int{10, 20, 30, 40, 50}

	tests := []struct {
		index    int
		expected []int
	}{
		{0, []int{20, 30, 40, 50}},
		{2, []int{10, 20, 40, 50}},
		{4, []int{10, 20, 30, 40}},
		{-1, []int{10, 20, 30, 40, 50}},
	}

	for _, tt := range tests {
		result := removeElement(copySlice(input), tt.index)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("removeElement(%v, %d) = %v; want %v", input, tt.index, result, tt.expected)
		}
	}
}
