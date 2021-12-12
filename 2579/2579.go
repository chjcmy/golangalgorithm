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

	arr, dp := make([]int, N+1), make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscanln(reader, &arr[i])
	}
	dp[1] = arr[1]
	if N >= 2 {
		dp[2] = arr[1] + arr[2]
	}
	for i := 3; i <= N; i++ {
		dp[i] = max(dp[i-2], dp[i-3]+arr[i-1]) + arr[i]
	}
	fmt.Fprint(writer, dp[N])
	writer.Flush()
}

func max(first, second int) int {
	if first >= second {
		return first
	}
	return second
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var reader *bufio.Reader = bufio.NewReader(os.Stdin)
//var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
//
//func main() {
//	var dp []int
//	var count int
//	fmt.Fscanln(reader, &count)
//	stairs := make([]int, 300)
//	for i := 0; i < count; i++ {
//		var num int
//		fmt.Fscanln(reader, &num)
//		stairs[i] = num
//	}
//	dp = append(dp, stairs[0])
//	dp = append(dp, Max(stairs[0] + stairs[1], stairs[0]))
//	dp = append(dp, Max(stairs[0]+stairs[2], stairs[1]+stairs[2]))
//	for i := 3; i < count; i++{
//			dp = append(dp, Max(dp[i-2]+stairs[i], stairs[i-1]+stairs[i]+dp[i-3]))
//	}
//	fmt.Fprintln(writer, dp[len(dp)-1])
//	writer.Flush()
//}
//
//func Max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
