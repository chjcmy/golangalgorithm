package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a *int, b *int, result *int, w *sync.WaitGroup) {
	defer w.Done()
	time.Sleep(time.Second)

	*result = *a + *b

}

func main() {
	var num1, num2 int = 10, 5
	var result int

	wait := sync.WaitGroup{}

	wait.Add(1)

	go add(&num1, &num2, &result, &wait)

	wait.Wait()

	fmt.Println(result)
}
