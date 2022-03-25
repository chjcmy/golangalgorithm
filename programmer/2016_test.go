package programmer

import (
	"fmt"
	"testing"
)

func Test2016(t *testing.T) {
	fmt.Println(Sol2016(5, 24))
}

func Sol2016(a int, b int) string {

	var total int

	w := []string{"FRI", "SAT", "SUN", "MON", "TUE", "WED", "THU"}

	m := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for i := 0; i < a-1; i++ {
		total += m[i]
	}

	total += b - 1

	return w[total%7]
}
