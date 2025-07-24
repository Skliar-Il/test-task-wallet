package main

import "fmt"

func MergeChannels(channels ...chan interface{}) chan interface{} {
	out := make(chan interface{}, len(channels)/2)
	go func() {
		done := make(chan struct{})
		for _, ch := range channels {
			go func(c chan interface{}) {
				for v := range c {
					out <- v
				}
				done <- struct{}{}
			}(ch)
		}

		for i := 0; i < len(channels); i++ {
			<-done
		}
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan interface{}, 5)
	ch2 := make(chan interface{}, 5)

	for i := 0; i < 5; i++ {
		ch1 <- i
		ch2 <- i + 5
	}
	close(ch1)
	close(ch2)

	merged := MergeChannels(ch1, ch2)

	for v := range merged {
		fmt.Println(v)
	}
}
