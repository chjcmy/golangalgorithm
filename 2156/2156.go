package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wines []int
	dp    []int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int

	fmt.Fscanln(reader, &N)

	wines = make([]int, N)
	dp = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &wines[i])
	}
	if N == 1 {
		fmt.Fprintln(writer, wines[0])
		writer.Flush()
		return
	} else if N == 2 {
		fmt.Fprintln(writer, wines[0]+wines[1])
		writer.Flush()
		return
	}
	dp[0] = wines[0]
	dp[1] = dp[0] + wines[1]
	dp[2] = Max()

	for i := 3; i < N-1; i++ {
		dp[i] = Max(dp[i-2]+wines[i], wines[i-1]+wines[i])
	}
	fmt.Fprintln(writer, dp[N-2]+wines[N-1])
	writer.Flush()
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
