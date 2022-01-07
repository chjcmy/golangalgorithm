package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)

	go square2(&wg, ch)
	ch <- 9

	wg.Wait()
}

func square2(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)

	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
