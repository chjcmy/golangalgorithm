package programmer

import "testing"

func TestMeasure(t *testing.T) {
	println(MeaSureSol(13, 17))
}

func MeaSureSol(left, right int) (result int) {

	var Even, Odd int

	ReMaining := right - left

	SellNum := make([][]int, ReMaining+1)

	for i := 0; i < len(SellNum); i++ {
		SellNum[i] = make([]int, 2)
	}

	for i := 0; i < len(SellNum); i++ {
		SellNum[i][0] = left + i
	}

	for i := 0; i < len(SellNum); i++ {
		for j := SellNum[i][0]; j > 0; j-- {
			if SellNum[i][0]%j == 0 {
				SellNum[i][1]++
			}
		}
		if SellNum[i][1]%2 == 0 {
			Even += SellNum[i][0]
		} else {
			Odd -= SellNum[i][0]
		}
	}

	return Even + Odd
}
