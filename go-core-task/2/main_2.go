package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateOriginalSlice() []int {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = r.Int()
	}
	return slice
}

func evenSlice(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, element int) []int {
	return append(slice, element)
}

func copySlice(slice []int) []int {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return copySlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("index out of range, returning original slice")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := generateOriginalSlice()
	fmt.Println("Original slice:", originalSlice)

	evanSliceRes := evenSlice(originalSlice)
	fmt.Println("Even numbers slice:", evanSliceRes)

	updatedSlice := addElements(originalSlice, 77)
	fmt.Println("After adding 77:", updatedSlice)

	copiedSlice := copySlice(originalSlice)
	originalSlice[0] = 999
	fmt.Println("Copied slice:", copiedSlice)
	fmt.Println("Modified original:", originalSlice)

	removedSlice := removeElement(originalSlice, 3)
	fmt.Println("After removing index 3:", removedSlice)
}
