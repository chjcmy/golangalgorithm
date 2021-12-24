package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	maps [][]graph
)

type graph struct {
	u int
	v int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, M, Start int

	fmt.Fscanf(reader, "%d %d\n", &N, &M)

	maps = make([][]graph, M+1)

	fmt.Fscanln(reader, Start)

	for i := 0; i <= M; i++ {
		var u, v, w int
		fmt.Fscanf(reader, "%d %d %d\n", &u, &v, &w)

		maps[u] = append(maps[u], graph{v, w})
	}

	for i := 0; i < M; i++ {
		fmt.Fprintln(writer, maps[i])
	}

	//for i := 0; i < M; i++ {
	//
	//}
	writer.Flush()
}
