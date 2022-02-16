package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, count int
	fmt.Fscanln(reader, &N)

	dp := [100][10]int{}
	for i := 1; i <= 9; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < N; i++ {
		dp[i][0] = dp[i-1][1]
		dp[i][9] = dp[i-1][8]
		for j := 1; j <= 8; j++ {
			dp[i][j] = (dp[i-1][j-1] + dp[i-1][j+1]) % 1000000000
		}
	}

	fmt.Print("")

	for i := 0; i < 10; i++ {
		count += dp[N-1][i]
	}
	fmt.Fprintln(writer, count%1000000000)
	writer.Flush()
}
