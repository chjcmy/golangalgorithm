package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go num(&wg)

	wg.Wait()
	fmt.Println("world")
}

func num(wg *sync.WaitGroup) {
	fmt.Println("hellow")
	wg.Done()
	return
}
