package main

import "testing"

func Test_NumberRandGenerator(t *testing.T) {
	out := make(chan int)
	count := 10
	go NumberRandGenerator(out, count)

	var counter int
	for range out {
		counter += 1
	}

	if counter != count {
		t.Errorf("expected %d numbers, got %d", count, counter)
	}
}
