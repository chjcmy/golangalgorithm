package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	fmt.Fscanln(reader, &N)

	dp := make([][]int, 101)
	for i := range dp {
		dp[i] = make([]int, 11)
	}
	writer.Flush()
}
