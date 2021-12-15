package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {

	var N int
	var stars, dp []int

	fmt.Fscanln(reader, &N)

	stars = make([]int, N)
	dp = make([]int, N)

	for i := 0; i < N; i++ {
		var P int
		fmt.Fscan(reader, &P)
		stars[i] = P
		dp[i] = 1
	}

	for i := 1; i < N; i++ {
		for j := 0; j <= i; j++ {
			if stars[j] < stars[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	man := dp[0]
	for i := 1; i < N; i++ {
		if man < dp[i] {
			man = dp[i]
		}
	}

	fmt.Fprintln(writer, man)
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
