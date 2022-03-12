package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

var (
	dp [][][]int
)

var Lines [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var a, b, c int
	dp := make([][][]int, 21)
	for i := range dp {
		dp[i] = make([][]int, 21)
		for j := range dp[i] {
			dp[i][j] = make([]int, 21)
		}
	}
	for {
		fmt.Fscanln(reader, &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, w(dp, a, b, c))
	}
	writer.Flush()
}
func w(dp [][][]int, a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	} else if a > 20 || b > 20 || c > 20 {
		return w(dp, 20, 20, 20)
	}
	if dp[a][b][c] != 0 {
		return dp[a][b][c]
	} else if a < b && b < c {
		dp[a][b][c] = w(dp, a, b, c-1) + w(dp, a, b-1, c-1) - w(dp, a, b-1, c)
	} else {
		dp[a][b][c] = w(dp, a-1, b, c) + w(dp, a-1, b-1, c) + w(dp, a-1, b, c-1) - w(dp, a-1, b-1, c-1)
	}
	return dp[a][b][c]
}
