package programmer

import (
	"sort"
	"testing"
)

func TestKnum(t *testing.T) {
	println(knumsol([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}))
}

func knumsol(array []int, commands [][]int) []int {

	var ret []int
	for _, cmd := range commands {
		slice := append([]int{}, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(slice)
		ret = append(ret, slice[cmd[2]-1])
	}
	return ret
}
