package programmer

import "testing"

func TestInsider(t *testing.T) {
	println(insidersol([]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}))
}

func insidersol(a []int, b []int) (result int) {

	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}

	return
}
