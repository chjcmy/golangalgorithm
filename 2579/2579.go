package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var (
	dp    [][]int
	count int
)

func main() {

	fmt.Fscanln(reader, &count)
	dp = make([][]int, count+1)
	dp[0] = append(dp[0], 0)
	for i := 1; i < count+1; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		dp[i] = append(dp[i], num)
	}
	for i := 0; i < count-3; {
		z := Max(i)
		dp[i+z][0] = dp[i][0] + dp[i+z][0]
		i += z

	}
	fmt.Fprintln(writer, Maxing(dp[count-1][0], dp[count-2][0]))
	writer.Flush()
}

func Max(a int) int {
	x := dp[a][0] + dp[a+1][0] + dp[a+3][0]
	y := dp[a][0] + dp[a+2][0] + dp[a+3][0]
	if x > y {
		return 1
	}
	return 2
}
func Maxing(x, y int) int {
	if x > y {
		return dp[count][0] + x
	}
	return dp[count][0] + y
}
