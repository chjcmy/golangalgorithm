package _Stesting

import (
	"strconv"
	"testing"
)

func TestTest2(t *testing.T) {
	println(test2Sol("82195", "64723"))
	println(test2Sol("00000000000000000000", "91919191919191919191"))
}

var minresult int

func test2Sol(p string, s string) int {

	for i := 0; i < len(p); i++ {
		pronum, _ := strconv.Atoi(string(p[i]))
		ansnum, _ := strconv.Atoi(string(s[i]))
		minresult += MaxSol(pronum, ansnum)
	}

	return minresult
}

func MaxSol(max, ans int) int {
	var now1 int

	if max > ans {
		if max-ans < 5 {
			now1 = max - ans
		} else {
			now1 = ans + 10 - max
		}
	}

	if ans > max {
		if ans-max < 5 {
			now1 = ans - max
		} else {
			now1 = max + 10 - ans
		}
	}
	return now1
}
