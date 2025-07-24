package main

import "testing"

func Test_UpgradeNumber(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	go UpgradeNumber(in, out)

	go func() {
		for i := 0; i < 10; i++ {
			in <- uint8(i)
		}
		close(in)
	}()

	expected := []float64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	for i := range expected {
		if value := <-out; value != expected[i] {
			t.Errorf("expected %f but got %f", expected[i], value)
		}
	}
}
