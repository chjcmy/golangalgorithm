package programmer

import "testing"

type numbool struct {
	absolutes []int
	signs     []bool
}

func TestUmYang(t *testing.T) {
	num1 := numbool{
		absolutes: []int{4, 7, 12},
		signs:     []bool{true, false, true},
	}

	println(UmYangsol(num1.absolutes, num1.signs))
}

func UmYangsol(absolutes []int, signs []bool) (result int) {

	for i := 0; i < len(absolutes); i++ {
		if signs[i] {
			result += absolutes[i]
		} else {
			result -= absolutes[i]
		}
	}

	return
}
