package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	dp = make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	var R, G, B int
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &R, &G, &B)
		dp[i][0], dp[i][1], dp[i][2] = R, G, B
	}
	for i := 1; i < N; i++ {
		dp[i][0] += min(dp[i-1][1], dp[i-1][2])
		dp[i][1] += min(dp[i-1][0], dp[i-1][2])
		dp[i][2] += min(dp[i-1][0], dp[i-1][1])
	}

	fmt.Fprintln(writer, min(dp[N-1][0], dp[N-1][1], dp[N-1][2]))

	writer.Flush()
}

func min(arr ...int) int {
	min := 9223372036854775807
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}
