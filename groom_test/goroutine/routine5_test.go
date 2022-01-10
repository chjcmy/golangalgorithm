package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRoutine5(t *testing.T) {

	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
