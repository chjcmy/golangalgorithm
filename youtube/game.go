package main

import (
	"fmt"
)

func main() {

	var a, b int
	var dir int
	var arr [2]int
	//var compass = [4][2]int{{-1, 0}, {0, 1}, {0, -1}, {}}
	//var visit [][]int
	fmt.Scanf("%d %d\n", &a, &b)
	m := make([][]int, a)
	for i := 0; i < b; i++ {
		m[i] = make([]int, b)
	}

	fmt.Scanf("%d %d %d\n", &arr[0], &arr[1], dir)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Scanf("%d", &m[i][j])
		}
	}

	fmt.Print(m)

}
