package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var N, Max, Min int

	var bytonic []int

	var dp []int

	fmt.Fscanln(reader, &N)

	bytonic = make([]int, N)

	dp = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &bytonic[i])

		if Max < bytonic[i] {
			Max = bytonic[i]
		}
	}

	Min = bytonic[0]

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
				ultraMax = bytonic[up]
				dp[i]++
			}
			ultraMax = bytonic[up]
		}
		for down := i + 1; down < N; down++ {
			if ultraMin > bytonic[down] {
				ultraMin = bytonic[down]
				dp[i]++
			}
		}
		if dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}
	fmt.Fprintln(writer, dp[N-1]+1)
	writer.Flush()
}
