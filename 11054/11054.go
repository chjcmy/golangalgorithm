package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

type dp struct {
	min int
	max int
}

func main() {
	var N int

	var bytonic []int

	fmt.Fscanln(reader, &N)

	bytonic = make([]int, N)
	deep := make([]dp, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &bytonic[i])
		deep[i].max, deep[i].min = bytonic[i], bytonic[i]
	}

	for i := 0; i < N; i++ {
		M := 0
		max := bytonic[i] - 1
		min := bytonic[i] - 1
		for up := 0; up < i; up++ {
			if bytonic[up] < max && max < bytonic[i] {
				max = bytonic[up]
				M++
			}
		}
		for down := i; down < N; down++ {
			if bytonic[down] > min && min < bytonic[i] {
				min = bytonic[down]
				M++
			}
		}
		dp[i] = M
	}

}
