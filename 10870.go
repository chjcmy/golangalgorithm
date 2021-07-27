package main

import (
	"bufio"
	"os"
	"strconv"
)

func fact2(n int64) int64 {

	b1 := int64(1)
	b2 := int64(1)
	b3 := int64(0)

	if n == 0 {
		b3 = 0
	} else if n <= 2 {
		b3 = 1
	} else {
		for i := int64(2); i < n; i++ {
			b3 = b1 + b2
			b1 = b2
			b2 = b3
		}
	}
	return b3
}

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	buff, _, _ := br.ReadLine()
	n, _ := strconv.ParseInt(string(buff), 10, 64)

	_, _ = bw.WriteString(strconv.FormatInt(fact2(n), 10))
	_ = bw.Flush()
}
