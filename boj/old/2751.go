package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)

func printf(f string, a ...interface{}) { _, _ = fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { _, _ = fmt.Fscanf(reader, f, a...) }

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int

	scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d\n", &arr[i])
	}

	sort.IntSlice(arr).Sort()
	for i := 0; i < n; i++ {
		printf("%d\n", arr[i])
	}
}
