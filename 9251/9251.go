package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var A, B []rune

	var S string

	var dp [][]int

	fmt.Fscanln(reader, &S)

	A = append(A, 0)

	A = append(A, []rune(S)...)

	dp = make([][]int, len(A))

	fmt.Fscanln(reader, &S)

	B = append(B, 0)

	B = append(B, []rune(S)...)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			dp[i] = make([]int, len(B))
		}
	}

	for i := 1; i < len(A); i++ {
		for j := i; j < len(B); j++ {
			if A[i] == B[j] {
				dp[i][j]++
			}
			dp[i][j] += dp[i][j-1]
		}
		if dp[i-1][len(B)-1] > dp[i][len(B)-1] {
			dp[i][len(B)-1] = dp[i-1][len(B)-1]
		}
	}

	fmt.Fprintln(writer, dp[len(A)-1][len(B)-1])

	writer.Flush()

}
