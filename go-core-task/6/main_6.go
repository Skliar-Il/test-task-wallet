package main

import (
	"fmt"
	"math/rand"
	"time"
)

func NumberRandGenerator(out chan<- int, count int) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := 0; i < count; i++ {
		out <- r.Int()
	}
	close(out)
}

func main() {
	out := make(chan int)
	go NumberRandGenerator(out, 10)

	for num := range out {
		fmt.Println(num)
	}
}
