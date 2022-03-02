package programmer

import (
	"fmt"
	"testing"
)

type loto struct {
	lotos   []int
	win_num []int
}

func TestLotto(t *testing.T) {

	loto1 := loto{
		lotos:   []int{44, 1, 0, 0, 31, 25},
		win_num: []int{31, 10, 45, 1, 6, 19},
	}

	loto2 := loto{
		lotos:   []int{0, 0, 0, 0, 0, 0},
		win_num: []int{38, 19, 20, 40, 15, 25},
	}

	loto3 := loto{
		lotos:   []int{45, 4, 35, 20, 3, 9},
		win_num: []int{20, 9, 3, 45, 4, 35},
	}

	fmt.Println(lotoSol(loto1.lotos, loto1.win_num))
	fmt.Println(lotoSol(loto2.lotos, loto2.win_num))
	fmt.Println(lotoSol(loto3.lotos, loto3.win_num))
}

func lotoSol(lottos []int, win_nums []int) []int {

	results := make([]int, 2)

	for i := 0; i < 2; i++ {
		var result int
		for j := 0; j < len(lottos); j++ {
			if i == 0 && lottos[j] == 0 {
				result++
				continue
			}
			if i == 1 && lottos[j] == 0 {
				continue
			}
			for k := 0; k < len(win_nums); k++ {
				if lottos[j] == win_nums[k] {
					result++
				}
			}
		}
		results[i] = score(result)
	}

	return results
}

func score(N int) (record int) {
	switch N {
	case 2:
		record = 5
	case 3:
		record = 4
	case 4:
		record = 3
	case 5:
		record = 2
	case 6:
		record = 1
	default:
		record = 6
	}
	return
}
