package main

import (
	"bufio"
	"os"
	"strconv"
)

func fact(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	buff, _, _ := br.ReadLine()
	n, _ := strconv.ParseInt(string(buff), 10, 64)

	bw.WriteString(strconv.FormatInt(fact(n), 10)) //nolint:errcheck
	bw.Flush()
}
