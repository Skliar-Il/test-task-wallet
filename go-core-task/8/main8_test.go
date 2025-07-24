package main

import (
	"slices"
	"testing"
)

func Test_CustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()
	result := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result = append(result, i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		if !slices.Contains(result, i) {
			t.Errorf("not expected %d in result, result: %v", i, result)
		}
	}
}
