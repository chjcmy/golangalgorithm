package closer

import (
	"fmt"
	"testing"
)

func next() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func TestCloser(t *testing.T) {
	nextInt := next()
	heyInt := next()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(heyInt())
	fmt.Println(heyInt())
	fmt.Println(nextInt())

	newInt := next()
	fmt.Println(newInt())
}
