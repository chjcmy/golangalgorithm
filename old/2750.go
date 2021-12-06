package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, elem int

	Fscanln(r, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		Fscanln(r, &elem)
		arr[i] = elem
	}
	sort.Ints(arr)
	for i := 0; i < n; i++ {
		Fprintln(w, arr[i])
	}
}
