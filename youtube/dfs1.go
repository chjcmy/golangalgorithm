package main

import "fmt"

var visited = []bool{false, false, false, false, false, false, false, false, false}
var graph = [][]int{
	{},
	{2, 3, 8},
	{1, 7},
	{1, 4, 5},
	{3, 5},
	{3, 4},
	{7},
	{2, 6, 8},
	{1, 7},
}

func main() {
	dfs(1)
}

func dfs(x int) {
	visited[x] = true
	fmt.Printf("%d ", x)
	for i := 0; i < len(graph[x]); i++ {
		y := graph[x][i]
		if !visited[y] {
			dfs(y)
		}
	}
}
