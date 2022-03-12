package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph [][]int
	out   int
	cnt   int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N int
	root := -1
	fmt.Fscanln(reader, &N)

	graph = make([][]int, N)

	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(reader, &a)
		if a == -1 {
			root = i
		} else {
			graph[a] = append(graph[a], i)
		}
	}

	fmt.Fscan(reader, &out)

	if out == root {
		fmt.Fprintln(writer, "0")
		writer.Flush()
		return
	}

	dfs(graph, root, out)

	fmt.Fprintln(writer, cnt)

	writer.Flush()
}

func dfs(graph [][]int, curr int, exc int) {
	isLeaf := 1
	for _, next := range graph[curr] {
		if next == exc {
			continue
		}
		isLeaf = 0
		dfs(graph, next, exc)
	}
	cnt += isLeaf
}
