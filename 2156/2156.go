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

	wines = make([]int, N+1)
	dp = make([]int, N+1)
	for i := 1; i < N+1; i++ {
		fmt.Fscanln(reader, &wines[i])
	}
	if N == 1 {
		fmt.Fprintln(writer, wines[1])
		writer.Flush()
		return
	} else if N == 2 {
		fmt.Fprintln(writer, wines[1]+wines[2])
		writer.Flush()
		return
	}
	dp[1] = wines[1]
	dp[2] = wines[1] + wines[2]

	for i := 3; i <= N; i++ {
		dp[i] = Max(dp[i-3]+wines[i-1]+wines[i], dp[i-1], dp[i-2]+wines[i])
	}
	fmt.Fprintln(writer, dp[N])
	writer.Flush()
}

func Max(arr ...int) int {
	var min int
	for _, i := range arr {
		if min < i {
			min = i
		}
	}
	return min
}
