package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

var a1, b1 int
var visited1 []bool
var arr1 []int
var w1 = bufio.NewWriter(os.Stdout)

func nnm2(idx int, cnt int) {
	if cnt == b1 {
		for _, v := range arr1 {
			w1.WriteString(strconv.Itoa(v) + " ")
		}
		w1.WriteString("\n")
		return
	}

	for i := idx + 1; i <= a1; i++ {
		if visited1[i] {
			continue
		}
		visited1[i] = true
		arr1[cnt] = i
		nnm2(i, cnt+1)
		visited1[i] = false
	}
}

func main() {
	Scan(&a1, &b1)

	visited1 = make([]bool, a1+1)
	arr1 = make([]int, b1)

	nnm2(0, 0)
	defer w1.Flush()
}
