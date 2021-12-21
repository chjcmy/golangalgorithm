package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, SUM int

	var grid []int

	fmt.Fscanln(reader, &N)

	grid = make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	sort.Ints(grid)

	for i := 1; i <= N; i++ {
		grid[i] = grid[i] + grid[i-1]
		SUM = grid[i] + SUM
	}

	fmt.Fprintln(writer, SUM)

	writer.Flush()
}
