package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type Line struct {
	from int
	to   int
}

type Lines []Line

func (l Lines) Len() int           { return len(l) }
func (l Lines) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l Lines) Less(i, j int) bool { return l[i].from < l[j].from }

func main() {
	var n int
	scanf("%d\n", &n)
	lines := make([]Line, n)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		var f, t int
		scanf("%d %d\n", &f, &t)
		lines[i] = Line{from: f, to: t}
		dp[i] = 1
	}
	sort.Sort(Lines(lines))

	max := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if lines[i].to > lines[j].to && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	printf("%d\n", n-max)
	writer.Flush()
}
