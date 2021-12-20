package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B = "0", "0"

	var C string

	var dp [][]int

	fmt.Fscanln(reader, &C)

	A += C

	fmt.Fscanln(reader, &C)

	B += C

	alen := len(A)

	blen := len(B)

	dp = make([][]int, alen)

	for i := 0; i < alen; i++ {
		for j := 0; j < blen; j++ {
			dp[i] = make([]int, blen)
		}
	}

	for i := 1; i < alen; i++ {
		for j := 1; j < blen; j++ {
			if A[i:i+1] == B[j:j+1] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}

	fmt.Fprintln(writer, dp[alen-1][blen-1])

	writer.Flush()

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
