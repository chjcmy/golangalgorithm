package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var N, Max int

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

	for i := 0; i < N; i++ {
		min := bytonic[0]
		if min > bytonic[i] {
			break
		}

	}

	for i := 1; i < N; i++ {

		if bytonic[i] < Max {
			dp[i] = dp[i-1]
			continue
		}

		var ultraMax int
		ultraMin := bytonic[i]
		for up := 0; up < i; up++ {
			if ultraMax < bytonic[up] {
				ultraMax = bytonic[up]
				dp[i]++
				continue
			}
			ultraMax = bytonic[up]
		}
		for down := i; down < N; down++ {
			if ultraMin < bytonic[down] {
				break
			}
			ultraMin = bytonic[down]
			dp[i]++
		}
		if dp[i-1] > dp[i] {
			dp[i] = dp[i-1]
		}
	}
	fmt.Fprintln(writer, dp[N-1])
	writer.Flush()
}
