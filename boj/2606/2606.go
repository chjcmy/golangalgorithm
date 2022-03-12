package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited []bool
	result  int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanln(reader, &N)

	graph = make([][]int, N+1)

	for i := 0; i <= N; i++ {
		graph[i] = make([]int, N+1)
	}

	visited = make([]bool, N+1)

	var count int
	fmt.Fscanln(reader, &count)

	for i := 1; i <= count; i++ {

		var com1, com2 int

		fmt.Fscanf(reader, "%d %d\n", &com1, &com2)

		graph[com1][com2] = 1
		graph[com1][com2] = 1
	}

	deepworm(1)

	fmt.Fprintln(writer, result-1)

}

func deepworm(start int) {

	visited[start] = true

	result += 1

	for i := 0; i < len(graph[start]); i++ {
		if graph[start][i] == 1 && !visited[i] {
			deepworm(i)
		}
	}
}
