package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	func main() {
		reader := bufio.NewReader(os.Stdin)
		writer := bufio.NewWriter(os.Stdout)

		var N int

		fmt.Fscanln(reader, &N)

		for i := 1; i <= N; i++ {
			fmt.Fscan(reader, &grid[i])
		}

		fmt.Fprintln(writer, grid[0])

		writer.Flush()
	}

}
