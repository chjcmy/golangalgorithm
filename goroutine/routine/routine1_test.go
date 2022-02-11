package routine

import (
	"fmt"
	"sync"
	"testing"
)

func TestRoutine1(t *testing.T) {
	var memoryAccess sync.Mutex

	var value int

	go func() {
		memoryAccess.Lock()

		value++

		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()

	if value == 0 {
		fmt.Printf("value is %d\t", value)
	} else {
		fmt.Println(value)
	}
}
