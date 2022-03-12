package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var (
	gun []int
	dp  []int
)

func main() {
	var bullets int
	max := 0
	fmt.Fscanln(reader, &bullets)
	gun = make([]int, bullets)
	for bullet := 0; bullet < bullets; bullet++ {
		fmt.Fscanln(reader, &gun[bullet])
		if max < gun[bullet] {
			max = gun[bullet]
		}
	}
	if max == 0 {
		fmt.Fprintf(writer, "%d", 1)
	}
	dp = make([]int, max+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 1
	for dp[max-1] == 0 {
		for con := 3; con < max+1; con++ {
			dp[con] = dp[con-2] + dp[con-3]
		}
	}
	for target := 0; target < len(gun); target++ {
		fmt.Fprintf(writer, "%d\n", dp[gun[target-1]])
	}
	writer.Flush()
}
