package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp [1000001]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + 1

		if i%2 == 0 {
			dp[i] = Min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = Min(dp[i], dp[i/3]+1)
		}
	}

	fmt.Fprintln(writer, dp[N])
	writer.Flush()
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
