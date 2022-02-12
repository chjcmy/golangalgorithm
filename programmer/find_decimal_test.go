package programmer

import (
	"testing"
)

func TestFind(t *testing.T) {

	number1 := "17"
	number2 := "011"

	finddemincal(number1)
	finddemincal(number2)
}

func finddemincal(pro string) {

	//isUsed := make([]int, len(pro))
	//
	//ans := 0
	//
	//num := make([]int, len(pro))
	//
	//for i, c := range pro {
	//	num[i] = int(c - '0')
	//}
	//
	//for i := 0; i < len(pro); i++ {
	//	if isUsed[i] {
	//
	//	}
	//}

}

func Binomial(s []string, r int) {

}

func demincal(val int) int {
	for i := 0; i*i <= val; i++ {
		if val%i == 0 {
			continue
		}
		return val
	}
	return -1
}
