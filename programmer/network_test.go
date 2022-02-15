package programmer

import (
	"fmt"
	"testing"
)

var (
	visited     []bool
	n           int
	computering [][]int
	result      int
)

func TestNetwork(t *testing.T) {

	n = 3

	computers := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}

	fmt.Println(netowrk(n, computers))
}

func netowrk(n int, computers [][]int) int {

	visited = make([]bool, n)

	computering = computers

	for i := 0; i < len(computers); i++ {
		if !visited[i] {
			dfs(i)
			result++
		}
	}

	return result
}

func dfs(node int) {
	visited[node] = true
	for i := 0; i < len(computering[node]); i++ {
		if !visited[i] && computering[node][i] == 1 && node != i {
			dfs(i)
		}
	}
}
