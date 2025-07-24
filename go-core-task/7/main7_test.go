package main

import (
	"slices"
	"testing"
)

func Test_MergeChannels(t *testing.T) {
	ch1 := make(chan interface{}, 5)
	ch2 := make(chan interface{}, 5)

	for i := 0; i < 5; i++ {
		ch1 <- i
		ch2 <- i + 5
	}
	close(ch1)
	close(ch2)

	merged := MergeChannels(ch1, ch2)

	var result []int
	for v := range merged {
		result = append(result, v.(int))
	}

	for i := 0; i < 5; i++ {
		if !slices.Contains(result, i) || !slices.Contains(result, i+5) {
			t.Errorf("expected values %d and %d in merged channel, got %v", i, i+5, result)
		}
	}
}
