package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	var stairs, dp []int

	fmt.Fscanln(reader, &N)
	stairs = make([]int, N)
	dp = make([]int, N)
	for stair := 0; stair < N; stair++ {
		fmt.Fscan(reader, &stairs[stair])
	}

	dp[0] = stairs[0]
	dp[1] = Max(dp[0], dp[0]+stairs[1])
	for i := 2; i < N-1; i++ {

		if -stairs[i-1] == stairs[i] && -stairs[i] < stairs[i+1] {
			dp[i] = dp[i-1] + stairs[i]
			dp[i+1] = dp[i]
			i++
			continue
		}
		if -stairs[i] > dp[i] || dp[i-1] == stairs[i] {
			dp[i] = dp[i-1]
			continue
		}

		if dp[i-1]+stairs[i] == dp[i] {
			dp[i+1] = dp[i]
			i++
			continue
		}

		if dp[i-1] > dp[i-1]+stairs[i]+stairs[i+1] {
			dp[i] = dp[i-2]
			continue
		}
		dp[i] = dp[i-1] + stairs[i]
	}
	fmt.Fprintln(writer, dp)
	writer.Flush()
}

func Max(arr ...int) int {
	var man int
	for _, i2 := range arr {
		if man < i2 {
			man = i2
		}
	}
	return man
}
