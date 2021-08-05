package main

import (
	"bufio"
	"fmt"
	"os"
)

var a, b int
var visited []bool
var arr []int
var writer = bufio.NewWriter(os.Stdout)

func result(cnt int) {
	if cnt == b {
		for v := 0; v < b; v++ {
			fmt.Fprintf(writer, "%d ", arr[v])
		}
		fmt.Fprintf(writer, "\n")
		return
	}

	for i := 1; i <= a; i++ {
		if !visited[i] {
			visited[i] = true
			arr[cnt] = i
			result(cnt + 1)
			visited[i] = false
		}
	}
}
func main() {
	fmt.Scan(&a, &b)

	visited = make([]bool, a+1)
	arr = make([]int, b)

	result(0)
	defer writer.Flush()
}
