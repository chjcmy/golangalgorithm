package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func square3(wg *sync.WaitGroup, ch chan int) {

	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("tick")
			time.Sleep(time.Second)
		case <-terminate:
			fmt.Println("terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}

	}
}

func TestRoutine4(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square3(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}
