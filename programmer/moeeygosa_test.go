package programmer

import (
	"testing"
)

var testval = [][]int{{1, 2, 3, 4, 5}, {2, 1, 2, 3, 2, 4, 2, 5}, {3, 3, 1, 1, 2, 2, 4, 4, 5, 5}}

func TestSolution(t *testing.T) {
	var pro = []int{1, 2, 3, 4, 5}
	sol(pro)
}

func sol(pro []int) (result []int) {

	var person = []int{0, 0, 0}

	maxint := 0

	for i := 0; i < len(person); i++ {

		for j := 0; j < len(pro); j++ {
			if pro[j] == testval[i][j%len(testval[i])] {
				person[i]++
			}
		}

		if maxint < person[i] {
			maxint = person[i]
		}
	}

	return maxmin(person, maxint)

}

func maxmin(arr []int, maxint int) (result []int) {

	for i, i2 := range arr {
		if i2 == maxint {
			result = append(result, i+1)
		}
	}

	return
}
