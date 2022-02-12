package programmer

import (
	"fmt"
	"testing"
)

func TestCarpet(t *testing.T) {

	fmt.Println(solution(10, 2))
	fmt.Println(solution(8, 1))
	fmt.Println(solution(24, 24))

}

type value struct {
	x int
	y int
}

func solution(brown int, yellow int) []int {

	values := new(value)

	for i := 1; i <= yellow; i++ {
		sero := i
		garo := yellow / sero

		if garo*sero == yellow {
			if brown == 2*(garo+sero)+4 {
				values.x = garo
				values.y = sero
				break
			}
		}
	}

	return []int{values.x + 2, values.y + 2}
}
