package closer

import (
	"fmt"
	"testing"
)

func calCoin(to, i int) func() int {
	coin := 0
	return func() int {
		coin += to * i //클로저
		return coin
	}
}

func TestCoin(t *testing.T) {

	coin10 := 1
	coin50 := 4
	coin100 := 5
	coin500 := 6

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Printf("잘못된 입력입니다.")
		return
	}

	add10 := calCoin(10, coin10)
	add50 := calCoin(50, coin50)
	add100 := calCoin(100, coin100)
	add500 := calCoin(500, coin500)

	totalmoney := add10() + add50() + add100() + add500()

	fmt.Println(totalmoney)
}
