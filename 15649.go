package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var a, b int
var visited []bool
var arr []int
var w = bufio.NewWriter(os.Stdout)

func result(cnt int) {
	if cnt == b {
		for _, v := range arr {
			w.WriteString(strconv.Itoa(v) + " ")
		}
		w.WriteString("\n")
		return
	}

	for i := 1; i <= a; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		arr[cnt] = i
		result(cnt + 1)
		visited[i] = false
	}
}

func main() {
	Scan(&a, &b)

	visited = make([]bool, a+1)
	arr = make([]int, b)

	result(0)
	defer w.Flush()
}
