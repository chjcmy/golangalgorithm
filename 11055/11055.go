package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var N, SN int
	var stars, dp []int

	fmt.Fscanln(reader, &N)

	stars = make([]int, N)
	dp = make([]int, N)

	for i := 0; i < N; i++ {
		var P int
		fmt.Fscan(reader, &P)
		stars[i] = P
	}

	dp[0] = stars[0]

	for i := 0; i < N; i++ {
		var M int
		for j := 0; j < i; j++ {
			if stars[j] < stars[i] {
				if M < dp[j] {
					M = dp[j]
				}
			}
		}
		dp[i] = M + stars[i]
		if SN < dp[i] {
			SN = dp[i]
		}
	}

	fmt.Fprintln(writer, SN)
	writer.Flush()
}
