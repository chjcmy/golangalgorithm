package Calc

import (
	"fmt"
	"math/rand"
	"testing"
)

var ac = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}
var bc = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}
var cc = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}
var dc = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}
var ec = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}
var ic = []int{rand.Intn(20),rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20),rand.Intn(20)}

func TestAdd(t *testing.T) {
	cal(10, ac)
	cal(10, bc)
	cal(10, cc)
	cal(10, dc)
	cal(10, ec)
	cal(10, ic)
}

func cal(a int, b []int) {
	var Max, Min int

	bytonic := b

	var dp []int

	N := a

	dp = make([]int, N)

	for i := 0; i < N; i++ {
		if Max < bytonic[i] {
			Max = bytonic[i]
		}
	}

	Min = bytonic[0]

	for i := 0; i < N; i++ {
		if Min < bytonic[i] {
			break
		}
		Min = bytonic[i]
		dp[0]++
	}

	for i := 1; i < N; i++ {

		ultraMax := bytonic[0]
		ultraMin := bytonic[i]
		for up := 1; up <= i; up++ {
			if ultraMax < bytonic[up] {
				dp[i]++
			}
			ultraMax = bytonic[up]
		}
		for down := i+1; down < N; down++ {
			if ultraMin > bytonic[down] {
				dp[i]++
			}
			ultraMin = bytonic[down]
		}
		if dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}
	fmt.Println(bytonic)
	fmt.Println(dp[N-1])
}
