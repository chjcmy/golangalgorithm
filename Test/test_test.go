package Test

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

var result int

func TestName(t *testing.T) {

	have := "PM 11:59:59"
	fmt.Println(try(have, 2))
}

func try(times string, n int) int {

	t1 := strings.Split(times, " ")

	t1 = append(t1, t1[0])

	t1 = t1[1:]

	t2 := t1[0] + t1[1]

	t3, _ := time.Parse("03:04:05PM", t2)

	for i := 0; i < n; i++ {
		t3 = t3.Add(time.Second)
	}

	fmt.Println(t3.Format("01:02:03"))

	return result
}
