package main

import "fmt"

func UpgradeNumber(in <-chan uint8, out chan<- float64) {
	for value := range in {
		out <- float64(value) * float64(value)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	go UpgradeNumber(in, out)
	go func() {
		for i := 0; i < 10; i++ {
			in <- uint8(i)
		}
		close(in)
	}()

	for value := range out {
		fmt.Println(value)
	}
}
