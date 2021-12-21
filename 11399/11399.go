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

	var N int

	fmt.Fscanln(reader, &N)

	var grid []int

	grid = make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Fscan(reader, &grid[i])
	}

	sort.Ints(grid)

	for i := 1; i <= N; i++ {
		grid[i] = grid[i] + grid[i-1]
		grid[0] = grid[i] + grid[0]
	}

	fmt.Fprintln(writer, grid[0])

	writer.Flush()
}
