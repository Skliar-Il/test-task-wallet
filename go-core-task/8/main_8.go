package main

import (
	"fmt"
	"sync/atomic"
)

type CustomWaitGroup struct {
	count int32
	done  chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		count: 0,
		done:  make(chan struct{}),
	}
}

func (w *CustomWaitGroup) Add(delta int) {
	if delta == 0 {
		return
	}
	if delta < 0 {
		panic("negative delta for CustomWaitGroup not allowed")
	}
	atomic.AddInt32(&w.count, int32(delta))
}

func (w *CustomWaitGroup) Done() {
	newVal := atomic.AddInt32(&w.count, -1)
	if newVal == 0 {
		close(w.done)
		return
	}
	if newVal < 0 {
		panic("negative count for CustomWaitGroup not allowed")
	}
}

func (w *CustomWaitGroup) Wait() {
	<-w.done
}

func main() {
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

	for _, v := range result {
		fmt.Println(v)
	}
}
