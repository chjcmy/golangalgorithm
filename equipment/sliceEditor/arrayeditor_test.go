package sliceEditor

import (
	"sort"
	"testing"
)

func TestArrayEditor(t *testing.T) {
	ArrayEditor([]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5}, {4, 4}, {1, 7}})
}

func ArrayEditor(array []int, commands [][]int) {

	for _, cmd := range commands {
		slice := append([]int{}, array[cmd[0]-1:cmd[1]]...)
		sort.Ints(slice)
	}
}
